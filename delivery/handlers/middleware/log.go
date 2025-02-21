package middleware

import (
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/itmrchow/microservice-gateway/delivery/response/writer"
	eErrs "github.com/itmrchow/microservice-gateway/entities/errors"
	mlog "github.com/itmrchow/microservice-gateway/entities/log"
	mCtx "github.com/itmrchow/microservice-gateway/infrastructure/util/context"
	mHttp "github.com/itmrchow/microservice-gateway/infrastructure/util/http"
)

// ApiLogHandler: 記錄req , resp info
func ApiLogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			rw = &writer.ResponseWriter{
				ResponseWriter: w,
			}
		)

		next.ServeHTTP(rw, r)

		event := getLogEvent(rw.StatusCode)

		logReq(event, r)
		logResp(event, rw)

		event.Send()
	})
}

// logReq: 記錄req info
func logReq(event *zerolog.Event, r *http.Request) {
	var (
		url     = r.URL.Path
		method  = r.Method
		ip      = mHttp.GetIP(r)
		traceID = mCtx.GetTraceID(r.Context())
	)

	// request info
	event.
		Str("type", "API").
		Str("trace_id", traceID).
		Str("url", url).
		Str("method", method).
		Str("ip", ip)

	switch method {
	case http.MethodGet, http.MethodDelete:
		event.Str("query_params", r.URL.RawQuery)
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			event.Err(err)
		}
		event.Str("request_body", string(body))
	}
}

// logResp: 記錄resp info
func logResp(event *zerolog.Event, rw *writer.ResponseWriter) {

	event.
		Int("status_code", rw.StatusCode).
		Str("response_body", string(rw.Data))
}

// getLogEvent: 取得log event , 根據statusCode 決定log level
func getLogEvent(statusCode int) *zerolog.Event {
	switch {
	case statusCode >= 500:
		return mlog.Err(errors.New(string(eErrs.SystemUnavailableErrCode)))
	case statusCode >= 400:
		return mlog.Warn()
	default:
		return mlog.Info()
	}
}
