# go-load
Docker repo to play with Golang and fire off some connections at end points

# Install

1. [Install go](https://golang.org/doc/install)
2. `go get github.com/dailyherold/go-load`
3. You can now run the built binary with `$GOPATH/bin/go-load`

# Use

`$GOPATH/bin/go-load [flags] target`

# Command line arguments

Currently the only supported flag is `-interview`, as this repo first started as an interview take home test! The flag is not required, but will apply the following logic to `target`:

1. Hit `target`, wait for one second, log to stdout.
2. Hit `target`, wait for two seconds, log to stdout.
3. Hit `target`, wait for four seconds, log to stdout.

If `target` is "https://httpbin.org/delay/3", the first two calls should timeout, and last call should succeed.
