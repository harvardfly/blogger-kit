package transport

import (
	"context"

	"github.com/gin-gonic/gin"

	"pkg.zpf.com/golang/blogger-kit/internal/app/article/endpoint"
	"pkg.zpf.com/golang/blogger-kit/internal/app/article/utils"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/middlewares/auth"
	httpware "pkg.zpf.com/golang/blogger-kit/internal/pkg/middlewares/http"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/utils/httputil"

	"go.uber.org/zap"

	kithttp "github.com/go-kit/kit/transport/http"
)

// NewHTTPHandler http handler use gin
func NewHTTPHandler(ctx context.Context, endpoints *endpoint.ArticleEndpoints, logger *zap.Logger) *gin.Engine {
	// 初始化gin
	r := httputil.NewRouter(ctx.Value("ginMod").(string), logger)
	r.Use(httpware.CommonMiddleware())
	options := httputil.ServerOptions(logger)
	// 初始化路由
	e := r.Group("/api/v1")
	e.Use(httpware.AccessControl())
	e.POST("/article", func(ctx *gin.Context) {
		kithttp.NewServer(
			endpoints.ArticleEndpoint,
			utils.DecodeArticleRequest,
			httputil.EncodeJSONResponse,
			options...,
		).ServeHTTP(ctx.Writer, ctx.Request)
		return
	})
	e.GET("/article", func(ctx *gin.Context) {
		kithttp.NewServer(
			auth.ValidJWTMiddleware(logger)(endpoints.GetArticleEndpoint),
			utils.DecodeArticleInfoRequest,
			httputil.EncodeJSONResponse,
			options...,
		).ServeHTTP(ctx.Writer, ctx.Request)
		return
	})

	return r
}
