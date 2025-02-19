package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	respDto "github.com/itmrchow/microservice-gateway/delivery/dto/resp"
	"github.com/itmrchow/microservice-gateway/delivery/handlers/middleware"
	"github.com/itmrchow/microservice-gateway/delivery/response/writer"
	eErrs "github.com/itmrchow/microservice-gateway/entities/errors"
)

// Public API
func RegisterPublicHandlers() *mux.Router {

	r := mux.NewRouter()

	// middleware
	// - TraceId
	// - recover panic
	// - log
	r.Use(middleware.Logger)

	v1 := r.PathPrefix("/v1").Subrouter()

	RegisterPublicHandlersV1(v1)

	return r
}

func RegisterPublicHandlersV1(r *mux.Router) {
	r.HandleFunc("/health", HealthHandler).Methods(http.MethodGet) // health check
	r.HandleFunc("/internal-error", InternalErrHandler).Methods(http.MethodGet)
	r.HandleFunc("/bad-request", BadReqHandler).Methods(http.MethodGet)
}

// TODO: 內容要移到usecase
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	resp := respDto.HealthDto{
		Message: "GATEWAY - HTTP is alive",
	}

	writer.SuccessResponseWriter(r, w, resp)
}

func InternalErrHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("internal server error")

	if err != nil {
		writer.InternalErrResponseWriter(r, w, err, nil)
	}

}

func BadReqHandler(w http.ResponseWriter, r *http.Request) {
	err := eErrs.NewBadRequestErr(eErrs.InvalidInputDataErrCode)
	writer.BadRequestErrResponseWriter(r, w, err, nil)
}
