package main

import (
	"context"
	"net/http"

	"github.com/be-ys/Genesis/internal/controllers"
	"github.com/be-ys/Genesis/internal/helpers"
	"github.com/be-ys/Genesis/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/zhashkevych/scheduler"
)

func main() {
	// Update targets right now
	services.UpdateTargets(context.Background())

	// Initialize worker to regularly pull data.
	worker := scheduler.NewScheduler()
	worker.Add(context.Background(), services.UpdateTargets, helpers.Configuration.CacheTimeDuration)

	// Start Web Server
	http.HandleFunc("/metrics", controllers.GetMetrics)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatalln(err)
	}
}
