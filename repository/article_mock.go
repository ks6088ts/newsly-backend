package repository

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ks6088ts/newsly-backend/graph/model"
)

// NewArticleMockRepository ...
func NewArticleMockRepository() ArticleRepository {
	return &articleMockRepository{
		db: map[string]*model.Article{},
	}
}

var _ ArticleRepository = (*articleMockRepository)(nil)

type articleMockRepository struct {
	sync.RWMutex
	db map[string]*model.Article
}

func (repo *articleMockRepository) Get(ctx context.Context, id string) (*model.Article, error) {
	repo.RLock()
	defer repo.RUnlock()

	article, ok := repo.db[id]
	if !ok {
		return nil, ErrNoSuchEntity
	}

	return article, nil
}

func (repo *articleMockRepository) GetAll(ctx context.Context) ([]*model.Article, error) {
	repo.RLock()
	defer repo.RUnlock()
	list := make([]*model.Article, 0, len(repo.db))

	for _, article := range repo.db {
		list = append(list, article)
	}

	// sort by CreatedAt
	sort.Slice(list, func(i, j int) bool {
		a := list[i]
		b := list[j]
		return a.CreatedAt.UnixNano() > b.CreatedAt.UnixNano()
	})

	return list, nil
}
func (repo *articleMockRepository) Create(ctx context.Context, article *model.Article) (*model.Article, error) {
	if article.ID != "" {
		return nil, ErrBadRequest
	}

	repo.Lock()
	defer repo.Unlock()

	article.ID = uuid.New().String()
	article.CreatedAt = time.Now()

	repo.db[article.ID] = article
	return article, nil
}

func (repo *articleMockRepository) Update(ctx context.Context, article *model.Article) (*model.Article, error) {
	if article.ID == "" {
		return nil, ErrBadRequest
	}

	repo.Lock()
	defer repo.Unlock()

	_, ok := repo.db[article.ID]
	if !ok {
		return nil, ErrNoSuchEntity
	}

	repo.db[article.ID] = article
	copyArticle := *article
	return &copyArticle, nil
}
