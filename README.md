# lotto24-api-test

API tests for Wikimedia public API

## Requirements

### For local run required:

* Golang 1.21+ [installed with proper GOHOME environment variable setted](https://go.dev/doc/install)

### for run in container:   

* Docker

## How to run

Open folder with project and run one of the following

### Local run

```
go test
```

### Build docker image and run in container:

```
docker build -t lotto24 .
docker run -it lotto24
```

### Run in docker container without building image:

```
docker run -it -v  .:/app -w /app golang:1.21-alpine go test
```

### See runs in CI

Just open github-actions page and see [all CI runs](https://github.com/temagi/lotto24-api-test/actions)


## Idea and implementation

