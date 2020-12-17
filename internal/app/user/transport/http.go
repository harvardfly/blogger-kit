package transport

import (
	"blogger-kit/internal/app/user/endpoint"
	"blogger-kit/internal/app/user/utils"
	"blogger-kit/internal/pkg/log"
	"context"
	"net/http"

	"go.uber.org/zap"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHttpHandler http handler use mux
func NewHttpHandler(ctx context.Context, endpoints *endpoint.UserEndpoints, logger *zap.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(log.NewZapLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(utils.EncodeError),
	}

	r.Methods("POST").Path("/register").Handler(kithttp.NewServer(
		endpoints.RegisterEndpoint,
		utils.DecodeRegisterRequest,
		utils.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/login").Handler(kithttp.NewServer(
		endpoints.LoginEndpoint,
		utils.DecodeLoginRequest,
		utils.EncodeJSONResponse,
		options...,
	))

	return r
}
