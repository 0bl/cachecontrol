package cachecontrol

import (
	"fmt"
	"strings"
	"time"
)

type CacheControl struct {
	maxAge          *time.Duration
	maxStale        *time.Duration
	minFresh        *time.Duration
	noCache         bool
	noStore         bool
	noTransform     bool
	onlyIfCached    bool
	public          bool
	private         bool
	mustRevalidate  bool
	proxyRevalidate bool
	sMaxAge         *time.Duration
}

// New method returns a new instance of the CacheControl struct.
func New() *CacheControl {
	return &CacheControl{}
}

// MaxAge method sets the maximum age of a cached response in seconds.
func (c *CacheControl) MaxAge(seconds int) *CacheControl {
	duration := time.Duration(seconds) * time.Second
	c.maxAge = &duration
	return c
}

// MaxStale method sets the maximum time a cached response can be used past its expiration time in seconds.
func (c *CacheControl) MaxStale(seconds int) *CacheControl {
	duration := time.Duration(seconds) * time.Second
	c.maxStale = &duration
	return c
}

// MinFresh method sets the minimum time a response can be considered fresh in seconds.
func (c *CacheControl) MinFresh(seconds int) *CacheControl {
	duration := time.Duration(seconds) * time.Second
	c.minFresh = &duration
	return c
}

// NoCache method indicates that the response cannot be served from a cache and must be revalidated with the origin server.
func (c *CacheControl) NoCache() *CacheControl {
	c.noCache = true
	return c
}

// NoStore method indicates that a cache must not store the response or any part of it.
func (c *CacheControl) NoStore() *CacheControl {
	c.noStore = true
	return c
}

// NoTransform indicates that the response should not be transformed (e.g., compressed, transcoded) by a proxy.
func (c *CacheControl) NoTransform() *CacheControl {
	c.noTransform = true
	return c
}

// OnlyIfCached method indicates that the client only wants a cached response, and not a new one.
func (c *CacheControl) OnlyIfCached() *CacheControl {
	c.onlyIfCached = true
	return c
}

// Public method indicates that the response may be cached by any cache, even if it is normally non-cacheable.
func (c *CacheControl) Public() *CacheControl {
	c.public = true
	return c
}

// Private method indicates that the response is intended for a single user and should not be cached by shared caches.
func (c *CacheControl) Private() *CacheControl {
	c.private = true
	return c
}

// MustRevalidate method indicates that a cache must revalidate the response with the origin server before using a stale cached response.
func (c *CacheControl) MustRevalidate() *CacheControl {
	c.mustRevalidate = true
	return c
}

// ProxyRevalidate method indicates that a cache shared by multiple users must revalidate the response with the origin server before using a stale cached response.
func (c *CacheControl) ProxyRevalidate() *CacheControl {
	c.proxyRevalidate = true
	return c
}

// SMaxAge method sets the maximum age of a cached response in seconds for shared caches.
func (c *CacheControl) SMaxAge(seconds int) *CacheControl {
	duration := time.Duration(seconds) * time.Second
	c.sMaxAge = &duration
	return c
}

func (c *CacheControl) Header() string {
	values := []string{}

	if c.noCache {
		values = append(values, "no-cache")
	}

	if c.noStore {
		values = append(values, "no-store")
	}

	if c.noTransform {
		values = append(values, "no-transform")
	}

	if c.onlyIfCached {
		values = append(values, "only-if-cached")
	}

	if c.public {
		values = append(values, "public")
	}

	if c.private {
		values = append(values, "private")
	}

	if c.mustRevalidate {
		values = append(values, "must-revalidate")
	}

	if c.proxyRevalidate {
		values = append(values, "proxy-revalidate")
	}

	if c.maxAge != nil {
		values = append(values, fmt.Sprintf("max-age=%d", int64(*c.maxAge/time.Second)))
	}

	if c.maxStale != nil {
		values = append(values, fmt.Sprintf("max-stale=%d", int64(*c.maxStale/time.Second)))
	}

	if c.minFresh != nil {
		values = append(values, fmt.Sprintf("min-fresh=%d", int64(*c.minFresh/time.Second)))
	}

	if c.sMaxAge != nil {
		values = append(values, fmt.Sprintf("s-maxage=%d", int64(*c.sMaxAge/time.Second)))
	}

	return strings.Join(values, ", ")
}
