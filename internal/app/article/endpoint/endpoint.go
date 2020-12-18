package endpoint

import (
	"blogger-kit/internal/app/article/service"
	"blogger-kit/internal/pkg/requests"
	"context"

	"github.com/go-kit/kit/endpoint"
)

// ArticleEndpoints 文章模块Endpoints
type ArticleEndpoints struct {
	ArticleEndpoint     endpoint.Endpoint
	GetArticleEndpoint  endpoint.Endpoint
	ArticleEditEndpoint endpoint.Endpoint
	ArticleDelEndpoint  endpoint.Endpoint
}

// MakeArticleEndpoint 创建文章Endpoint
func MakeArticleEndpoint(articleService service.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		articleReq := req.(*requests.Article)
		articleInfo, err := articleService.Article(ctx, articleReq)
		return articleInfo, err
	}
}

// MakeGetArticleEndpoint 查询文章Endpoint
func MakeGetArticleEndpoint(articleService service.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		articleReq := req.(*requests.ArticleInfo)
		articleInfo, err := articleService.GetArticle(articleReq)
		if err != nil {
			return nil, err
		}
		return articleInfo, nil
	}
}
