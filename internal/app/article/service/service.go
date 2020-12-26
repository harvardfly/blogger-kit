package service

import (
	"context"
	"errors"

	"pkg.zpf.com/golang/blogger-kit/internal/app/article/dao"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/models"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/requests"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/responses"

	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// ArticleService 定义article service
type ArticleService interface {
	Article(ctx context.Context, req *requests.Article) (responses.Article, error)
	GetArticle(req *requests.ArticleInfo) (responses.ArticleRes, error)
	ArticleEdit(req *requests.ArticleEdit) error
	ArticleDel(req *requests.ArticleInfo) error
}

// ArticleServiceImpl 初始默认的
type ArticleServiceImpl struct {
	articleDao dao.ArticleDao
	logger     *zap.Logger
}

// NewArticleServiceImpl 初始化
func NewArticleServiceImpl(articleDao dao.ArticleDao, logger *zap.Logger) ArticleService {
	return &ArticleServiceImpl{
		articleDao: articleDao,
		logger:     logger.With(zap.String("type", "NewArticleServiceImpl")),
	}
}

// Article 发表文章
func (s *ArticleServiceImpl) Article(ctx context.Context, req *requests.Article) (responses.Article, error) {
	var result responses.Article
	span, _ := opentracing.StartSpanFromContext(ctx, "article.service")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer span.Finish()
	cateInfo := s.articleDao.GetCategory(req.CategoryID)
	if cateInfo.ID == 0 {
		s.logger.Error("categoryId 不存在", zap.Any("category_id", req.CategoryID))
		return result, gorm.ErrRecordNotFound
	}

	res, err := s.articleDao.Article(ctx, req)

	if err != nil {
		s.logger.Error("发表文章失败", zap.Error(err))
		return result, errors.New("发表文章失败")
	}

	result = responses.Article{
		ID:    res.ID,
		Title: res.Title,
	}
	return result, err
}

// GetArticle 获取文章详情
func (s *ArticleServiceImpl) GetArticle(req *requests.ArticleInfo) (responses.ArticleRes, error) {
	var result responses.ArticleRes
	info := s.articleDao.GetArticle(req.ID)
	result = responses.ArticleRes{
		ID:        info.ID,
		Title:     info.Title,
		Summary:   info.Summary,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
		Category: responses.Category{
			ID:        info.Category.ID,
			Name:      info.Category.Name,
			Num:       info.Category.Num,
			CreatedAt: info.Category.CreatedAt,
			UpdatedAt: info.Category.UpdatedAt,
		},
	}
	return result, nil
}

// ArticleEdit 更新文章  全量更新
func (s *ArticleServiceImpl) ArticleEdit(req *requests.ArticleEdit) error {
	articleInfo := s.articleDao.GetArticle(req.ID)
	cateInfo := s.articleDao.GetCategory(req.CategoryID)
	if articleInfo.ID == 0 || cateInfo.ID == 0 {
		s.logger.Error("articleId或categoryId不存在", zap.Any("article_id", req.ID), zap.Any("category_id", req.CategoryID))
		return gorm.ErrRecordNotFound
	}

	article := &models.Article{
		CategoryID: req.CategoryID,
		Summary:    req.Summary,
		Title:      req.Title,
	}
	return s.articleDao.ArticleEdit(article)
}

// ArticleDel 删除文章
func (s *ArticleServiceImpl) ArticleDel(req *requests.ArticleInfo) error {
	info := s.articleDao.GetArticle(req.ID)
	if info.ID == 0 {
		s.logger.Error("articleId不存在", zap.Any("article_id", req.ID))
		return gorm.ErrRecordNotFound
	}
	return s.articleDao.ArticleDel(info.ID)
}
