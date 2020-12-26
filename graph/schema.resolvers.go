package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/ks6088ts/newsly-backend/graph/generated"
	"github.com/ks6088ts/newsly-backend/graph/model"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.CreateArticleInput) (*model.Article, error) {
	return r.ArticleRepository.Create(ctx, &model.Article{
		URL:       input.URL,
		Title:     input.Title,
		CreatedAt: time.Now(),
		Source: &model.Source{
			ID:   input.SourceID,
			Name: "user " + input.SourceID,
		},
	})
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	return r.ArticleRepository.GetAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
