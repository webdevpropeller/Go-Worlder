package localcache

import "time"

// Cache ...
type Cache interface {
	Set(k string, x interface{}, d time.Duration)
	Get(k string) (interface{}, bool)
	DefaultExpiration() time.Duration
}
