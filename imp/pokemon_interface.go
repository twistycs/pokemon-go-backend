package imp

import "github.com/twistycs/pokemon-go-backend/models"

type PokemonService interface {
	GetAllPokemon() (pokemonList []models.Pokemon, err error)
	GetPokemonByPokemonName(pokemonName string) (pokemon models.Pokemon, err error)
}

type PokemonRepository interface {
	GetAllPokemon(u *[]models.Pokemon) (err error)
	GetPokemonByPokemonName(u *models.Pokemon, pokemonName string) (err error)
}
