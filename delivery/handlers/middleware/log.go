package middleware

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/itmrchow/microservice-gateway/util"
)

// Logger: 記錄req , resp info
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			logEvent = log.Info()

			url     = r.URL.Path
			method  = r.Method
			ip      = util.GetIP(r)
			traceID = util.GetTraceID(r.Context())
		)

		// request info
		logEvent.
			Str("type", "API").
			Str("trace_id", traceID).
			Str("url", url).
			Str("method", method).
			Str("ip", ip)

		logQueryParams(logEvent, r)

		next.ServeHTTP(w, r)

		//response info

		logEvent.Send()
	})
}

func logQueryParams(logEvent *zerolog.Event, r *http.Request) {

}
