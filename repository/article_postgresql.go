package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	// postgresql
	_ "github.com/lib/pq"

	"github.com/ks6088ts/newsly-backend/graph/model"
)

type articlePostgresqlRepository struct {
	db *gorm.DB
}

// NewArticlePostgresqlRepository ...
func NewArticlePostgresqlRepository(dataSourceName string) (ArticleRepository, error) {
	db, err := gorm.Open("postgres", dataSourceName)
	fmt.Println(dataSourceName)
	if err != nil {
		return nil, err
	}

	return &articlePostgresqlRepository{
		db: db,
	}, nil
}

func (repo *articlePostgresqlRepository) Get(ctx context.Context, id string) (*model.Article, error) {
	article := model.Article{}
	result := repo.db.First(&article, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// WIP: support source
	article.Source = &model.Source{
		ID:   "0",
		Name: "not supported",
	}
	return &article, result.Error
}

func (repo *articlePostgresqlRepository) GetAll(ctx context.Context) ([]*model.Article, error) {
	articles := []*model.Article{}
	result := repo.db.Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}

	// WIP: support source
	for _, article := range articles {
		article.Source = &model.Source{
			ID:   "0",
			Name: "not supported",
		}
	}
	return articles, nil
}

func (repo *articlePostgresqlRepository) Create(ctx context.Context, article *model.Article) (*model.Article, error) {
	result := repo.db.Create(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	return article, nil
}

func (repo *articlePostgresqlRepository) Update(ctx context.Context, article *model.Article) (*model.Article, error) {
	return nil, ErrNotImplemented
}
