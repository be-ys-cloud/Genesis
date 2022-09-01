# Genesis

Genesis is an open-source project, created by the Enterprise Architecture Team. His main aim is to provide a simple,
flexible system to ping various services, and export all these metrics in a readable format for Prometheus.

## How does it work?

The project is a Golang program, running a web server. The server exposes a `/metrics` endpoint on port 8080, who shows
all ping time and HTTP status.

Cache time and endpoints could be easily configured in the `config.json` file.

## How to execute it?

* Stand-alone / dev : `go mod download && go run ./cmd`
* Docker : `docker build --no-cache -t="genesis:latest" . && docker run -d --name="genesis" -p 8080:8080 -v config.json:/app/config.json genesis:latest`


## Example of data

```
genesis_http_request_status{service="website",url="https://be-ys.com/"} 200
genesis_http_request_time{service="website",url="https://be-ys.com/"} 0.091947
genesis_http_request_status{service="website",url="https://www.be-ys.com/ecosysteme-be-ys"} 200
genesis_http_request_time{service="website",url="https://www.be-ys.com/ecosysteme-be-ys"} 0.087947
genesis_http_request_status{service="serviceA",url="https://serviceA.local/"} 200
genesis_http_request_time{service="serviceA",url="https://serviceA.local/"} 0.094947
```

## License

Project released under MIT License.