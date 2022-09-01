package services

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func fetch(service string, url string) string {
	start := time.Now()
	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(url)
	var responseCode = res.StatusCode
	if err != nil {
		responseCode = 504
		logrus.Warningf("URL %s for service %s is not reachable !", url, service)
	}

	val := strconv.FormatFloat(time.Now().Sub(start).Seconds(), 'f', 6, 64)

	var writer bytes.Buffer
	_, _ = fmt.Fprintf(&writer, "genesis_http_request_status{service=\"%s\",url=\"%s\"} %d\n", service, url, responseCode)
	_, _ = fmt.Fprintf(&writer, "genesis_http_request_time{service=\"%s\",url=\"%s\"} %s\n", service, url, val)
	return writer.String()
}
