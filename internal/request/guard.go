package request

import "context"

func Active(ctx context.Context) bool {
	return ctx != nil && ctx.Err() == nil
}
