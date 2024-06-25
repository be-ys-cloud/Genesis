package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/be-ys/Genesis/internal/structures"
	"github.com/sirupsen/logrus"
)

func fetch(service string, url string, headers map[string]string, extract map[string]structures.ExtractData, format string) string {
	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Warningf("Error trying to request %s: %s", url, err.Error())
	}

	// Add headers to http request
	for header_key, header_value := range headers {
		req.Header.Add(header_key, header_value)
	}
	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	var responseCode int
	if err != nil {
		responseCode = 504
		logrus.Warningf("URL %s for service %s is not reachable !", url, service)
	} else {
		responseCode = res.StatusCode
	}

	val := strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64)

	var writer bytes.Buffer
	_, _ = fmt.Fprintf(&writer, "genesis_http_request_status{service=\"%s\",url=\"%s\"} %d\n", service, url, responseCode)
	_, _ = fmt.Fprintf(&writer, "genesis_http_request_time{service=\"%s\",url=\"%s\"} %s\n", service, url, val)

	if len(extract) > 0 && responseCode >= 200 && responseCode < 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			logrus.Warningf("Can't read body: %s", "")
			return writer.String()
		}
		extractData(&writer, service, url, body, extract, format)
	}
	return writer.String()
}

func extractData(writer *bytes.Buffer, service string, url string, body []byte, extract map[string]structures.ExtractData, format string) string {
	switch format {
	case "json":
		var result interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			logrus.Warningf("Error while parsing json: %s", err.Error())
			return writer.String()
		}
		for field, _ := range extract {
			fieldName := extract[field].FieldName
			value, found := findValueInJson(result, fieldName)
			if found {
				mapField := extract[field]
				if len(mapField.Values) == 0 {
					switch valueType := value.(type) {
					case float32:
					case float64:
						_, _ = fmt.Fprintf(writer, "genesis_http_request_field{service=\"%s\",url=\"%s\",field=\"%s\"} %f\n", service, url, field, value)
					default:
						logrus.Warningf("%s value is not a number (type: %s). Please add 'values' to your configuration", field, reflect.TypeOf(valueType).String())
					}
				} else {
					keyValue, exists := mapField.Values[value.(string)]
					if exists {
						_, _ = fmt.Fprintf(writer, "genesis_http_request_field{service=\"%s\",url=\"%s\",field=\"%s\"} %d\n", service, url, field, keyValue)
					} else {
						logrus.Warningf("values do not contains %s for %s", value.(string), fieldName)
					}

				}
			} else {
				logrus.Warningf("fieldname %s does not exist", fieldName)
			}
		}
		return writer.String()
	default:
		return writer.String()
	}
}

func findValueInJson(data interface{}, key string) (interface{}, bool) {
	keys := strings.Split(key, ".")
	for _, key := range keys {
		switch value := data.(type) {
		case map[string]interface{}:
			if v, found := value[key]; found {
				data = v
			} else {
				return nil, false
			}
		default:
			return nil, false
		}
	}
	return data, true
}
