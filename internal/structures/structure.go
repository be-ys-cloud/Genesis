package structures

import "time"

type Config struct {
	CacheTime         string        `json:"cache_time"`
	CacheTimeDuration time.Duration `json:"-"`
	Hosts             []ConfigHosts `json:"hosts"`
}

type ConfigHosts struct {
	Name      string     `json:"prometheusName"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Name        string                 `json:"name"`
	Headers     map[string]string      `json:"headers"`
	ExtractJson map[string]ExtractData `json:"extractJson"`
}

type ExtractData struct {
	FieldName string         `json:"fieldName"`
	Values    map[string]int `json:"values"`
}
