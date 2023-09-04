package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/GuilhermeSa/go-service-template/internal/clients"
	"github.com/GuilhermeSa/go-service-template/internal/handlers"
	"github.com/GuilhermeSa/go-service-template/internal/repositories"
	"github.com/go-chi/chi"
	"github.com/golobby/container/v3"
)

func Start(config Config) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := loadDependencies()
	if err != nil {
		panic(fmt.Errorf("loading dependencies: %w", err))
	}
	r := chi.NewRouter()
	r.Use(requestTracking)
	r.Use(structuredLogger)
	r.Mount("/v1", v1Router())
	slog.Info("starting server")
	http.ListenAndServe(":8080", r)
}

func v1Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/pokemon/nationaldex/{id}", handlers.GetPokemonByNationalDex)
	return r
}

func loadDependencies() error {
	err := container.Singleton(func() repositories.PokemonRepository {
		return clients.NewPokeAPIClient()
	})
	return err
}
