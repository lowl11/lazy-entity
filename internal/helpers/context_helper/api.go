package context_helper

import (
	"context"
	"time"
)

type CtxFunc = func(customTimeout ...time.Duration) (context.Context, func())

func GetContext(requestCtx context.Context, getCtxFunc CtxFunc, customTimeout ...time.Duration) (context.Context, func()) {
	if requestCtx == nil || requestCtx == context.Background() || requestCtx == context.TODO() {
		return getCtxFunc(customTimeout...)
	}

	return requestCtx, func() {}
}
