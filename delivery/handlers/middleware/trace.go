package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const TraceIDHeader = "X-Trace-ID"

type TraceIDKey struct{}

// Trace: 添加 trace id 到 context
func Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(TraceIDHeader)
		if traceID == "" {
			traceID = uuid.New().String()
			r.Header.Set(TraceIDHeader, traceID)
		}
		// save in context
		ctx := context.WithValue(r.Context(), TraceIDKey{}, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
