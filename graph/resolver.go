package graph

import "github.com/ks6088ts/newsly-backend/repository"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	ArticleRepository repository.ArticleRepository
}
