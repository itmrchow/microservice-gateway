package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Public API
func RegisterPublicHandlers() *mux.Router {

	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()

	RegisterPublicHandlersV1(v1)

	return r
}

func RegisterPublicHandlersV1(r *mux.Router) {
	r.HandleFunc("/health", HealthHandler).Methods(http.MethodGet) // health check
}

// TODO: 要移到usecase
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(" GATEWAY - HTTP is alive"))
	if err != nil {
		return
	}
}
