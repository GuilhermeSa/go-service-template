package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GuilhermeSa/go-service-template/internal/model"
)

const (
	pokeAPIURL = "https://pokeapi.co/api/v2/"
)

type PokeAPI struct {
	httpClient *http.Client
}

func NewPokeAPIClient() *PokeAPI {
	return &PokeAPI{
		httpClient: &http.Client{
			Timeout: time.Second,
		},
	}
}

func (p *PokeAPI) GetPokemonByNationalDex(id int) (model.Pokemon, error) {
	resp, err := p.httpClient.Get(fmt.Sprintf("%s/pokemon/%d", pokeAPIURL, id))
	if err != nil {
		return model.Pokemon{}, fmt.Errorf("requesting pokemon from pokeapi: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return model.Pokemon{}, fmt.Errorf("getting pokemon from pokeapi: %w", err)
	}
	var pokemon model.Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return model.Pokemon{}, fmt.Errorf("decoding pokemon from pokeapi: %w", err)
	}
	return pokemon, nil
}
