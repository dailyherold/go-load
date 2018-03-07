# go-load
Docker repo to play with Golang and fire off some connections at end points

# Install

1. [Install go](https://golang.org/doc/install)
2. `go get github.com/dailyherold/go-load`
3. You can now run the built binary with `$GOPATH/bin/go-load`

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
