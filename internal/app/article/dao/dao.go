package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/databases"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/models"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/requests"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/responses"

	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	"go.uber.org/zap"
)

// ArticleDao dao数据层方法 crud
type ArticleDao interface {
	Article(ctx context.Context, req *requests.Article) (models.Article, error)
	GetArticle(id int) models.Article
	ArticleEdit(article *models.Article) error
	ArticleDel(id int) error
	GetCategory(id int) models.Category
}

// ArticleDaoImpl 默认实现
type ArticleDaoImpl struct {
	logger *zap.Logger
}

// NewArticleDaoImpl 初始化
func NewArticleDaoImpl(logger *zap.Logger) ArticleDao {
	return &ArticleDaoImpl{
		logger: logger.With(zap.String("type", "NewArticleDaoImpl")),
	}
}

// withID 通过单个ID查询对应记录
func withID(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// withIds 通过一组ID查询对应记录
func withIds(ids []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	}
}

// GetArticle 获取文章及相应的类别信息
func (r *ArticleDaoImpl) GetArticle(id int) models.Article {
	var value models.Article
	databases.DB.Preload("Category").Model(&models.Article{}).Scopes(withID(id)).First(&value)
	return value
}

// GetCategory 根据类别获取类别对象信息
func (r *ArticleDaoImpl) GetCategory(id int) models.Category {
	var value models.Category
	databases.DB.Model(&models.Category{}).Scopes(withID(id)).First(&value)
	return value
}

// Article 发布文章
func (r *ArticleDaoImpl) Article(ctx context.Context, req *requests.Article) (models.Article, error) {
	var category models.Category
	value := models.Article{
		Title:      req.Title,
		Summary:    req.Summary,
		CategoryID: req.CategoryID,
	}
	span, _ := opentracing.StartSpanFromContext(ctx, "article.mysql")
	defer span.Finish()
	err := databases.DB.Model(&models.Article{}).Create(&value).Error
	if err != nil {
		return value, err
	}
	databases.DB.Model(&models.Category{}).Scopes(withID(value.CategoryID)).First(&category)
	esValue := requests.ArticleES{
		ID:      value.ID,
		Summary: value.Summary,
		Title:   value.Title,
		Category: responses.Category{
			ID:        value.CategoryID,
			Name:      category.Name,
			Num:       category.Num,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
		CreatedAt: value.CreatedAt,
		UpdatedAt: value.UpdatedAt,
	}
	fmt.Println(esValue)

	return value, err
}

// ArticleEdit 文章全量更新
func (r *ArticleDaoImpl) ArticleEdit(article *models.Article) error {
	err := databases.DB.Model(&models.Article{}).Scopes(withID(article.ID)).UpdateColumns(map[string]interface{}{
		"category_id": article.CategoryID,
		"summary":     article.Summary,
		"title":       article.Title,
		"updated_at":  time.Now(),
	}).Error
	if err != nil {
		r.logger.Error("更新文章失败", zap.Error(err))
		return errors.New("更新文章失败")
	}

	return nil
}

// ArticleDel 删除文章
func (r *ArticleDaoImpl) ArticleDel(id int) error {
	err := databases.DB.Scopes(withID(id)).Delete(&models.Article{}).Error
	if err != nil {
		r.logger.Error("删除文章失败", zap.Error(err))
		return errors.New("删除文章失败")
	}

	return nil
}
