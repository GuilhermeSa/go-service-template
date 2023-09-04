package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/GuilhermeSa/go-service-template/internal/repositories"
	"github.com/GuilhermeSa/go-service-template/internal/usecases"
	"github.com/go-chi/chi"
	"github.com/golobby/container/v3"
)

func GetPokemonByNationalDex(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var pokemonRepository repositories.PokemonRepository
	err = container.Resolve(&pokemonRepository)
	if err != nil {
		slog.Error("trying to resolve pokemon repository")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pokemon, err := usecases.GetPokemonByNationalDex(pokemonRepository, id)
	response, err := json.Marshal(pokemon)
	w.Write(response)
}
