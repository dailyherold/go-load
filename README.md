# go-load
Docker repo to play with Golang and fire off some connections at end points.

# Install

1. [Install go](https://golang.org/doc/install).
2. `go get github.com/dailyherold/go-load`.
3. You can now run the built binary with `$GOPATH/bin/go-load`.

# Use

`$GOPATH/bin/go-load -target=https://httpbin.org`

# Command line arguments

`$GOPATH/bin/go-load -h` will output help text for available arguments.

```
  -interview
        Boolean flag for running interview logic on target
  -target string
        Target URL for load testing
```

`-target` is required.

`-interview` is optional, as this repo first started as an interview take home test! It will apply the following logic to your target URL:

1. Hit target, wait for one second, log to stdout.
2. Hit target, wait for two seconds, log to stdout.
3. Hit target, wait for four seconds, log to stdout.

If `-target=https://httpbin.org/delay/3`, the first two calls timeout, and last call succeeds.

# Docker

A multi-stage Dockerfile allows a very small Docker image to be built with a compiled go-load binary.

## Build image

Build the image by running:

`docker build -t $IMAGE:$TAG .`

Example:

`docker build -t github.com/dailyherold/go-load:0.0.1 .`

## Run image

The image's `ENTRYPOINT` is `./load` which allows you to pass any go-load arguments in your docker run command like you would a local binary. `-h` is passed by default to display help text.

```
❯ docker run --rm github.com/dailyherold/go-load:0.0.1
Usage of ./load:
  -interview
        Boolean flag for running interview logic on target
  -target string
        Target URL for load testing
```

Example of passing a `-target` arg:

`docker run --rm github.com/dailyherold/go-load:0.0.1 -target=https://httpbin.org`

## Docker Hub

This repository is [linked](https://hub.docker.com/r/dailyherold/go-load/) to a public Docker Hub automated build service. Any merges to `master` will kick off the build process and push the built image to Docker Hub.

Pull the `latest` image with `docker pull dailyherold/go-load`.
