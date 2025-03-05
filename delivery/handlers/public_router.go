package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	accountV1 "github.com/itmrchow/microservice-proto/account/v1"

	respDto "github.com/itmrchow/microservice-gateway/delivery/dto/resp"
	"github.com/itmrchow/microservice-gateway/delivery/handlers/middleware"
	"github.com/itmrchow/microservice-gateway/delivery/response/writer"
	eErrs "github.com/itmrchow/microservice-gateway/entities/errors"
	"github.com/itmrchow/microservice-gateway/infrastructure/svc"
)

// Public API
func RegisterPublicHandlers() *mux.Router {

	r := mux.NewRouter()

	// middleware
	r.Use(middleware.Trace)
	// - recover panic
	r.Use(middleware.ApiLogHandler)

	v1 := r.PathPrefix("/v1").Subrouter()

	RegisterInternalHandlersV1(v1)

	return r
}

// Internal API
func RegisterInternalHandlers() *mux.Router {

	r := mux.NewRouter()

	// middleware
	r.Use(middleware.Trace)
	// - recover panic
	r.Use(middleware.ApiLogHandler)

	v1 := r.PathPrefix("/v1").Subrouter()

	RegisterInternalHandlersV1(v1)

	return r
}

func RegisterInternalHandlersV1(r *mux.Router) {
	r.HandleFunc("/health", HealthHandler).Methods(http.MethodGet) // health check
	r.HandleFunc("/internal-error", InternalErrHandler).Methods(http.MethodGet)
	r.HandleFunc("/bad-request", BadReqHandler).Methods(http.MethodGet)

	r.HandleFunc("/account/user", GetAccountUserV1).Methods(http.MethodGet)
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

func GetAccountUserV1(w http.ResponseWriter, r *http.Request) {

	// id := r.URL.Query().Get("id")
	// email := r.URL.Query().Get("email")

	// // TODO:驗證
	// if id == "" && email == "" {
	// 	err := eErrs.NewBadRequestErr(eErrs.InvalidInputDataErrCode)
	// 	writer.BadRequestErrResponseWriter(r, w, err, nil)
	// 	return
	// }

	userSvc, err := svc.NewAccountUserSvcV1()
	if err != nil {
		return
	}

	resp, err := userSvc.GetUser(r.Context(), &accountV1.GetUserRequest{
		Id: "123",
	})

	if err != nil {
		return
	}

	// DTO
	userResp := respDto.GetAccountUserV1Resp{}
	userResp.FromProto(resp)

	writer.SuccessResponseWriter(r, w, userResp)
}
