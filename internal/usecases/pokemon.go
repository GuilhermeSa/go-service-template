package usecases

import (
	"github.com/GuilhermeSa/go-service-template/internal/model"
	"github.com/GuilhermeSa/go-service-template/internal/repositories"
)

func GetPokemonByNationalDex(repository repositories.PokemonRepository, id int) (model.Pokemon, error) {
	return repository.GetPokemonByNationalDex(id)
}
