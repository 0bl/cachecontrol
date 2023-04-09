# cachecontrol [![Go](https://github.com/0bl/cachecontrol/actions/workflows/go.yml/badge.svg)](https://github.com/0bl/cachecontrol/actions/workflows/go.yml) [![Go Report](https://goreportcard.com/badge/github.com/0bl/cachecontrol)](https://goreportcard.com/report/github.com/0bl/cachecontrol) [![GoDoc](https://godoc.org/github.com/0bl/cachecontrol?status.svg)](https://godoc.org/github.com/0bl/cachecontrol)
cachecontrol is a go package for manipulating http cache control headers.

## installation
to use cachecontrol in your go project, simply run:
```sh
> go get github.com/0bl/cachecontrol
```

## usage
first, import the package:
```go
import "github.com/0bl/cachecontrol"
```
to create a new instance of `CacheControl`, use the `New()` method:
```go
cc := cachecontrol.New()
```
you can then use the available methods to set the cache control directives you need. for example:
```go
cc.MaxAge(3600)
cc.Private()
```
to retrieve the cache control header string, use the `Header()` method:
```go
header := cc.Header()
fmt.Println(header)
// Output: "max-age=3600, private"
```
