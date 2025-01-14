package middleware

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
)

func AddLoggerToContext(logger *zap.Logger) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = context.WithValue(ctx, CtxLoggerKey("logger"), logger)
		return next(ctx)
	}
}

// AddOperationToContext adds operation details to the context.
func Logger(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	l := ctx.Value(CtxLoggerKey("logger")).(*zap.Logger)

	start := time.Now()

	response := next(ctx)

	responseCtx := graphql.GetOperationContext(ctx)

	if l != nil {
		l.Info(
			"request",
			zap.String("content_type", responseCtx.Headers.Get("Content-Type")),
			zap.String("user_agent", responseCtx.Headers.Get("User-Agent")),
			zap.String("operation_name", responseCtx.Operation.Name),
			zap.Duration("duration", time.Since(start)),
		)
	}

	return response
}
