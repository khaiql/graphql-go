package ratelimiting

import "context"

type RateLimiter interface {
	SaveUsage(ctx context.Context, selection string, isLimitted bool)
	IsLimited(ctx context.Context, selection string) bool
}

type NoOp struct{}

func (n *NoOp) SaveUsage(ctx context.Context, selection string, isLimitted bool) {
}

func (n *NoOp) IsLimited(ctx context.Context, selection string) bool {
	return false
}
