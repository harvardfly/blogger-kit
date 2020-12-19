package transport

import (
	"blogger-kit/internal/app/user/endpoint"
	"blogger-kit/internal/app/user/utils"
	httpware "blogger-kit/internal/pkg/middlewares/http"
	"blogger-kit/internal/pkg/utils/httputil"
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// NewHttpHandler http handler use mux
func NewHttpHandler(ctx context.Context, endpoints *endpoint.UserEndpoints, logger *zap.Logger) http.Handler {
	r := mux.NewRouter()
	options := httputil.ServerOptions(logger)
	r.Use(httpware.AccessControl)
	r.Methods("POST").Path("/register").Handler(kithttp.NewServer(
		endpoints.RegisterEndpoint,
		utils.DecodeRegisterRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/login").Handler(kithttp.NewServer(
		endpoints.LoginEndpoint,
		utils.DecodeLoginRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/findByID").Handler(kithttp.NewServer(
		endpoints.FindByIDEndpoint,
		utils.DecodeFindByIDRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	return r
}
