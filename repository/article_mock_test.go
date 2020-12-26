package repository

import (
	"context"
	"testing"

	"github.com/ks6088ts/newsly-backend/graph/model"
)

func Test_ArticleRepository(t *testing.T) {
	ctx := context.Background()
	repo := NewArticleMockRepository()

	articleA, err := repo.Create(ctx, &model.Article{
		Title: "test A",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.Update(ctx, &model.Article{
		ID:    articleA.ID,
		Title: "test A!",
	})
	if err != nil {
		t.Fatal(err)
	}

	articleA2, err := repo.Get(ctx, articleA.ID)
	if err != nil {
		t.Fatal(err)
	}
	if v1, v2 := articleA.ID, articleA2.ID; v1 != v2 {
		t.Errorf("unexpected: %#v, %#v", v1, v2)
	}

	articleB, err := repo.Create(ctx, &model.Article{
		Title: "test B",
	})
	if err != nil {
		t.Fatal(err)
	}

	list, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if v := len(list); v != 2 {
		t.Fatalf("unexpected: %#v", v)
	}
	if v := list[0]; v.ID != articleB.ID {
		t.Fatalf("unexpected: %#v", v)
	}
	if v := list[1]; v.ID != articleA.ID {
		t.Fatalf("unexpected: %#v", v)
	}
}
