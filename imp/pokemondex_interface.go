package imp

import "github.com/twistycs/pokemon-go-backend/models"

type PokemonDexService interface {
	GetAllPokemonDex() (pokemonDexList []models.PokemonDex, err error)
	GetPokemonDexByPokemonDexName(pokemonDexName string) (pokemonDex models.PokemonDex, err error)
}

type PokemonDexRepository interface {
	GetAllPokemonDex(u *[]models.PokemonDex) (err error)
	GetPokemonDexByPokemonDexName(u *models.PokemonDex, pokemonDexName string) (err error)
}
