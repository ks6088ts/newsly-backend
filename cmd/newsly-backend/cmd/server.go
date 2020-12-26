/*
Copyright © 2020 ks6088ts <ks6088ts@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package cmd ...
package cmd

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/ks6088ts/newsly-backend/graph"
	"github.com/ks6088ts/newsly-backend/graph/generated"
	"github.com/ks6088ts/newsly-backend/repository"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func getArticleRepository() (repository.ArticleRepository, error) {
	db := os.Getenv("DB_TYPE")
	if db == "mock" {
		return repository.NewArticleMockRepository(), nil
	}
	if db == "postgresql" {
		return repository.NewArticlePostgresqlRepository(os.Getenv("DB_DATA_SOURCE_NAME"))
	}
	return nil, fmt.Errorf("%v is not supported", db)
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start GraphQL server",
	Long:  `start GraphQL server`,
	Run: func(cmd *cobra.Command, args []string) {
		router := chi.NewRouter()

		// ref. https://gqlgen.com/recipes/cors/
		// Add CORS middleware around every request
		// See https://github.com/rs/cors for full option listing
		router.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{os.Getenv("ALLOWED_ORIGIN")},
			AllowCredentials: true,
			Debug:            true,
		}).Handler)

		const defaultPort = "8080"

		port := os.Getenv("PORT")
		if port == "" {
			port = defaultPort
		}

		repo, err := getArticleRepository()
		if err != nil {
			fmt.Printf("failed to get repository: %v\n", err)
			os.Exit(1)
		}

		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
			ArticleRepository: repo,
		}}))

		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		router.Handle("/query", srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
