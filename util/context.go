package util

import "context"

// context key
type TraceIDKey struct{}

// Get func

// GetTraceID: 從 context 中獲取 trace ID
func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(TraceIDKey{}).(string)
	if !ok {
		return ""
	}
	return traceID
}
