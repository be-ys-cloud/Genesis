package helpers

import (
	"encoding/json"
	"github.com/be-ys/Genesis/internal/structures"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

var Configuration structures.Config

func init() {
	//Load configuration file
	file, err := os.Open("config.json")
	if err != nil {
		logrus.Fatalln("Unable to open config.json !")
	}

	fileContent, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(fileContent, &Configuration)
	_ = file.Close()

	Configuration.CacheTimeDuration, err = time.ParseDuration(Configuration.CacheTime)
	if err != nil {
		logrus.Warnf("Provided refresh time %s is invalid. Falling back to 15s.", Configuration.CacheTime)
		Configuration.CacheTimeDuration = time.Second * 15
	}
}
