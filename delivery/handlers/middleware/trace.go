package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	mCtx "github.com/itmrchow/microservice-gateway/infrastructure/util/context"
)

const TraceIDHeader = "X-Trace-ID"

// Trace: 添加 trace id 到 context
func Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(TraceIDHeader)
		if traceID == "" {
			traceID = uuid.New().String()
			r.Header.Set(TraceIDHeader, traceID)
		}
		// save in context
		ctx := context.WithValue(r.Context(), mCtx.TraceIDKey{}, traceID)
		w.Header().Set(TraceIDHeader, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
