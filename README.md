# time-service

A simple http server that gives you current time in RFC3339 format.

## How to use

```shell
$ go build main.go
$ ./main
```

### Example

```shell
$ curl http://localhost:3000
{"date_time":"2024-04-16T00:16:40+08:00"}
```

## Configuration

You can configure time-service using environment variable.

* `PORT`: the port that http server should listen to
