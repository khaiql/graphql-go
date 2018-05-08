package authorization

import "context"

type Authorizor interface {
	CanDo(ctx context.Context, selection string) bool
}

type NoOp struct{}

func (n *NoOp) CanDo(ctx context.Context, selection string) bool {
	return true
}
