/*
	Copyright (c) 2020 be|ys - MIT License
	For more informations, please refer to the LICENSE file.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var configuration Config

func main() {
	//Load configuration file
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Unable to open config.json !")
	}

	fileContent, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(fileContent, &configuration)
	_ = file.Close()

	//Start Web Server
	http.Handle("/metrics", cached(configuration.CacheTime, pingServices))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func pingServices(w http.ResponseWriter, r *http.Request) {
	resc := make(chan string)

	for _, host := range configuration.Hosts {
		for _, url := range host.Endpoints {
			go func(host string, url string) {
				body, _ := fetch(host, url)
				resc <- body
			}(host.Name, url)
		}
	}

	final := ""
	for i := 0; i < len(configuration.Hosts); i++ {
		for j := 0; j < len(configuration.Hosts[i].Endpoints); j++ {
			final = final + <-resc + "\n"
		}
	}

	_, err := fmt.Fprintf(w, final)

	if err != nil {
		fmt.Println(err)
	}
}

func fetch(host string, url string) (string, error) {
	start := time.Now()
	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(url)
	var responseCode = "504"
	if err == nil {
		responseCode = strconv.Itoa(res.StatusCode)
	}

	val := strconv.FormatFloat(time.Now().Sub(start).Seconds(), 'f', 6, 64)
	return "genesis_http_request_status{service=\"" + host + "\",url=\"" + url + "\"} " + responseCode + "\n" + "genesis_http_request_time{service=\"" + host + "\",url=\"" + url + "\"} " + val, nil
}
