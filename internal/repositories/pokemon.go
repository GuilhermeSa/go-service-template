package repositories

import "github.com/GuilhermeSa/go-service-template/internal/model"

type PokemonRepository interface {
	GetPokemonByNationalDex(id int) (model.Pokemon, error)
}
