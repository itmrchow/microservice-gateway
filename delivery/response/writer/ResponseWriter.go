package writer

import (
	"encoding/json"
	"net/http"

	"github.com/itmrchow/microservice-common/response"
	"github.com/rs/zerolog/log"

	"github.com/itmrchow/microservice-gateway/entities/errors"
	mCtx "github.com/itmrchow/microservice-gateway/infrastructure/util/context"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Data       []byte
}

func (r *ResponseWriter) Header() http.Header {
	return r.ResponseWriter.Header()
}

func (r *ResponseWriter) Write(p0 []byte) (int, error) {
	r.Data = p0
	return r.ResponseWriter.Write(p0)
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func SuccessResponseWriter(r *http.Request, w http.ResponseWriter, data any) {
	resp := response.SuccessResponse{}
	resp.Message = "success"
	resp.Data = data

	jsonData, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func InternalErrResponseWriter(r *http.Request, w http.ResponseWriter, err error, errData any) {
	// funcName, line := system.GetCaller(2)

	// TODO: print err log

	traceID := mCtx.GetTraceID(r.Context())

	log.Error().
		Str("trace_id", traceID).
		Str("func_name", "").
		Str("line", "").
		Err(err).
		Msg("internal server error")

	resp := response.FailedResponse{}
	resp.Message = string(errors.SystemUnavailableErrCode)
	resp.Error = "internal server error"
	resp.Data = errData

	jsonData, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func BadRequestErrResponseWriter(r *http.Request, w http.ResponseWriter, err errors.BadRequestErr, errData any) {
	resp := response.FailedResponse{}

	resp.Error = "bad request"
	resp.Message = err.Error()
	resp.Data = errData
	jsonData, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
