package gacos

import "sync"

type gacos struct {
	endPoint string
	cacheMd5 string
}

var (
	once sync.Once
	g    *gacos
)

func SingleGacos(endPoint string) *gacos {
	once.Do(func() {
		g = &gacos{endPoint: endPoint}
	})
	return g
}
