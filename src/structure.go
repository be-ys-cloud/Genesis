/*
	Copyright (c) 2020 be|ys - MIT License
	For more informations, please refer to the LICENSE file.
*/

package main

type Config struct {
	CacheTime string        `json:"cache_time"`
	Hosts     []ConfigHosts `json:"hosts"`
}

type ConfigHosts struct {
	Name      string   `json:"prometheusName"`
	Endpoints []string `json:"endpoints"`
}
