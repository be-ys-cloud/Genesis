package controllers

import (
	"github.com/be-ys/Genesis/internal/services"
	"net/http"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	_, _ = w.Write([]byte(services.Telemetry))
}
