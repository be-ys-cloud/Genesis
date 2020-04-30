# Genesis

Genesis is an open-source project, created by the Enterprise Architecture Team. His main aim is to provide a simple, flexible system to ping various services, and export all these metrics in a readable format for Prometheus.

## Why did you do that? There already some libraries for this
Indeed yes, for some frameworks or languages, there already are some alternatives. For exemple, Spring Actuator. But we easily can find some backwards to this kind of tools:
* You get metrics from the app, not from the outside: If your system is working but not accessible (eg. because of a network disruption), you will not be able to see it in your monitoring.
* For light systems (eg. a simple nginx who is serving static files), there is no actuators available.

## How does it work?
The project is a Golang program, running a web server. The server exposes a `/metrics` endpoint on port 8080, who shows all ping time and HTTP status.

Cache time and endpoints could be easily configured in the `config.json` file.

## How to execute it?
Simply run `go build src/`. A Quick-Setup is also available, through Docker. Simply download the Dockerfile in a folder, put your own config.json in the same folder, and build-it !

For docker :
```sh
docker build --no-cache -t="genesis:latest" ./genesis
docker run -d --name="genesis" -p 8080:8080 genesis:latest
```

## Example of data
```
genesis_http_request_status{service="website",url="https://be-ys.com/"} 200
genesis_http_request_time{service="website",url="https://be-ys.com/"} 0.091947
genesis_http_request_status{service="website",url="https://www.be-ys.com/ecosysteme-be-ys"} 200
genesis_http_request_time{service="website",url="https://www.be-ys.com/ecosysteme-be-ys"} 0.087947
genesis_http_request_status{service="serviceA",url="https://serviceA.local/"} 200
genesis_http_request_time{service="serviceA",url="https://serviceA.local/"} 0.094947
```

## Acknowledgement & License
* Thanks to [https://github.com/goenning/go-cache-demo](https://github.com/goenning/go-cache-demo) for the cache system.

Project released under MIT License.