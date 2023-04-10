# CacheControl [![Go](https://github.com/0bl/cachecontrol/actions/workflows/go.yml/badge.svg)](https://github.com/0bl/cachecontrol/actions/workflows/go.yml) [![Go Report](https://goreportcard.com/badge/github.com/0bl/cachecontrol)](https://goreportcard.com/report/github.com/0bl/cachecontrol) [![GoDoc](https://godoc.org/github.com/0bl/cachecontrol?status.svg)](https://godoc.org/github.com/0bl/cachecontrol)
CacheControl is a Go package for manipulating HTTP cache control headers.

## Installation
To use CacheControl in your Go project, simply run:
```
go get github.com/0bl/cachecontrol
```

## Usage
First, import the package:
```go
import "github.com/0bl/cachecontrol"
```
To create a new instance of `CacheControl`, use the `New()` method:
```go
cc := cachecontrol.New()
```
You can then use the available methods to set the cache control directives you need. For example:
```go
cc.MaxAge(3600)
cc.Private()
```
To retrieve the cache control header string, use the `Header()` method:
```go
header := cc.Header()
fmt.Println(header)
// Output: "max-age=3600, private"
```

## Available Methods
The following methods are available to set cache control directives:
- `MaxAge(seconds int) *CacheControl`: sets the maximum age of a cached response in seconds.
- `MaxStale(seconds int) *CacheControl`: sets the maximum time a cached response can be used past its expiration time in seconds.
- `MinFresh(seconds int) *CacheControl`: sets the minimum time a response can be considered fresh in seconds.
- `NoCache() *CacheControl`: indicates that the response cannot be served from a cache and must be revalidated with the origin server.
- `NoStore() *CacheControl`: indicates that a cache must not store the response or any part of it.
- `NoTransform() *CacheControl`: indicates that the response should not be transformed (e.g., compressed, transcoded) by a proxy.
- `OnlyIfCached() *CacheControl`: indicates that the client only wants a cached response, and not a new one.
- `Public() *CacheControl`: indicates that the response may be cached by any cache, even if it is normally non-cacheable.
- `Private() *CacheControl`: indicates that the response is intended for a single user and should not be cached by shared caches.
- `MustRevalidate() *CacheControl`: indicates that a cache must revalidate the response with the origin server before using a stale cached response.
- `ProxyRevalidate() *CacheControl`: indicates that a cache shared by multiple users must revalidate the response with the origin server before using a stale cached response.
- `SMaxAge(seconds int) *CacheControl`: sets the maximum age of a cached response in seconds for shared caches.
