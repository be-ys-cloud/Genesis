package structures

import "time"

type Config struct {
	CacheTime         string        `json:"cache_time"`
	CacheTimeDuration time.Duration `json:"-"`
	Hosts             []ConfigHosts `json:"hosts"`
}

type ConfigHosts struct {
	Name      string   `json:"prometheusName"`
	Endpoints []string `json:"endpoints"`
}
