package transport

import (
	"blogger-kit/internal/app/article/endpoint"
	"blogger-kit/internal/app/article/utils"
	"blogger-kit/internal/pkg/middlewares/auth"
	httpware "blogger-kit/internal/pkg/middlewares/http"
	"blogger-kit/internal/pkg/utils/httputil"
	"context"
	"net/http"

	"go.uber.org/zap"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPHandler http handler use mux
func NewHTTPHandler(ctx context.Context, endpoints *endpoint.ArticleEndpoints, logger *zap.Logger) http.Handler {
	r := mux.NewRouter()
	options := httputil.ServerOptions(logger)
	r.Use(httpware.AccessControl)
	r.Methods("POST").Path("/article").Handler(kithttp.NewServer(
		endpoints.ArticleEndpoint,
		utils.DecodeArticleRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/article").Handler(kithttp.NewServer(
		auth.ValidJWTMiddleware(logger)(endpoints.GetArticleEndpoint),
		utils.DecodeArticleInfoRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	return r
}
