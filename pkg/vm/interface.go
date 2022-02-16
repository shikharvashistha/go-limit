package vm

type vmRateLimiter interface {
	Limit()
}
