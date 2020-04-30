/*
	Copyright (c) 2020 be|ys - MIT License
	For more informations, please refer to the LICENSE file.

	Original code by goenning - MIT License. Please refer to the README file for a direct link.
*/

package main

import (
	"net/http"
	"net/http/httptest"
	"time"
)

var storage = NewStorage()

func cached(duration string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := storage.Get(r.RequestURI)
		if content != nil {
			_, _ = w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)

			w.WriteHeader(c.Code)
			content := c.Body.Bytes()

			if d, err := time.ParseDuration(duration); err == nil {
				storage.Set(r.RequestURI, content, d)
			}
			_, _ = w.Write(content)
		}
	})
}
