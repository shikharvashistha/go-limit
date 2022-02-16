package disk

type rateLimiter interface {
	Limit()
}
