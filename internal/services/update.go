package services

import (
	"bytes"
	"context"
	"github.com/be-ys/Genesis/internal/helpers"
	"sync"
)

func UpdateTargets(_ context.Context) {

	var writer bytes.Buffer
	var wg sync.WaitGroup

	for _, host := range helpers.Configuration.Hosts {
		for _, url := range host.Endpoints {
			wg.Add(1)
			go func(host string, url string) {
				writer.Write([]byte(fetch(host, url)))
				wg.Done()
			}(host.Name, url)
		}
	}

	wg.Wait()

	Telemetry = writer.String()
}
