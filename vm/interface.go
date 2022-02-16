package vm

type rateLimiter interface {
	Limit()
}
