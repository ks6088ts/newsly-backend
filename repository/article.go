package repository

import (
	"context"

	"github.com/ks6088ts/newsly-backend/graph/model"
)

// ArticleRepository ...
type ArticleRepository interface {
	Get(ctx context.Context, id string) (*model.Article, error)
	GetAll(ctx context.Context) ([]*model.Article, error)
	Create(ctx context.Context, article *model.Article) (*model.Article, error)
	Update(ctx context.Context, article *model.Article) (*model.Article, error)
}
