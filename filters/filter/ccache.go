package filter

import (
	"errors"

	"github.com/karlseguin/ccache"
	"time"
)

var cache *ccache.Cache

func init() {

	cache = ccache.New(ccache.Configure().MaxSize(100).ItemsToPrune(5))
}

func Set(key string, value string) {
	cache.Set(key, value, time.Second*10)
}

func Get(key string) (string, error) {
	item := cache.Get(key)
	if item != nil {
		if item.TTL().Seconds() > 0 {
			res, ok := item.Value().(string)
			if ok {
				return res, nil
			} else {
				return "", errors.New("cache item 类型错误")
			}

		}
	}
	return "", errors.New("cache item empty")
}
