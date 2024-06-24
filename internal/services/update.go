package services

import (
	"bytes"
	"context"
	"sync"

	"github.com/be-ys/Genesis/internal/helpers"
	"github.com/be-ys/Genesis/internal/structures"
)

func UpdateTargets(_ context.Context) {

	var writer bytes.Buffer
	var wg sync.WaitGroup

	for _, host := range helpers.Configuration.Hosts {
		for _, endpoint := range host.Endpoints {
			var extract map[string]structures.ExtractData
			var format string
			if endpoint.ExtractJson != nil {
				extract = endpoint.ExtractJson
				format = "json"
			}
			wg.Add(1)
			go func(host string, url string, headers map[string]string, extract map[string]structures.ExtractData, format string) {
				writer.Write([]byte(fetch(host, url, headers, extract, format)))
				wg.Done()
			}(host.Name, endpoint.Name, endpoint.Headers, extract, format)
		}
	}

	wg.Wait()

	Telemetry = writer.String()
}
