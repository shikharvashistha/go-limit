package disk

type diskRateLimiter interface {
	Limit()
}
