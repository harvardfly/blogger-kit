package transport

import (
	"context"

	"github.com/gin-gonic/gin"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/user/endpoint"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/user/utils"
	httpware "pkg.zpf.com/golang/kit-scaffold/internal/pkg/middlewares/http"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/utils/httputil"

	kithttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

// NewHTTPHandler http handler use mux
func NewHTTPHandler(ctx context.Context, endpoints *endpoint.UserEndpoints, logger *zap.Logger) *gin.Engine {
	// 初始化gin
	r := httputil.NewRouter(ctx.Value("ginMod").(string), logger)
	r.Use(httpware.CommonMiddleware())
	options := httputil.ServerOptions(logger)

	// 初始化路由
	e := r.Group("/api/v1")
	e.Use(httpware.AccessControl())
	e.POST("/register", func(ctx *gin.Context) {
		kithttp.NewServer(
			endpoints.RegisterEndpoint,
			utils.DecodeRegisterRequest,
			httputil.EncodeJSONResponse,
			options...,
		).ServeHTTP(ctx.Writer, ctx.Request)
		return
	})
	e.POST("/login", func(ctx *gin.Context) {
		kithttp.NewServer(
			endpoints.LoginEndpoint,
			utils.DecodeLoginRequest,
			httputil.EncodeJSONResponse,
			options...,
		).ServeHTTP(ctx.Writer, ctx.Request)
	})
	e.GET("/findByID", func(ctx *gin.Context) {
		kithttp.NewServer(
			endpoints.FindByIDEndpoint,
			utils.DecodeFindByIDRequest,
			httputil.EncodeJSONResponse,
			options...,
		).ServeHTTP(ctx.Writer, ctx.Request)
	})

	return r
}
