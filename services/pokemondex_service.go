package services

import (
	"fmt"

	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
)

type pokemonDexService struct {
	pokemonDexRepo imp.PokemonDexRepository
}

func NewPokemonDexService(repo imp.PokemonDexRepository) imp.PokemonDexService {
	return &pokemonDexService{pokemonDexRepo: repo}
}

func (pokemonDexService *pokemonDexService) GetAllPokemonDex() (pokemonDexs []models.PokemonDex, err error) {
	var pokemonDex []models.PokemonDex
	handle := pokemonDexService.pokemonDexRepo.GetAllPokemonDex(&pokemonDex)
	if handle != nil {
		fmt.Println("Error")
	}
	return pokemonDex, handle
}

func (pokemonDexService *pokemonDexService) GetPokemonDexByPokemonDexName(pokemonDexName string) (pokemonDexs models.PokemonDex, err error) {
	var pokemonDex models.PokemonDex
	handle := pokemonDexService.pokemonDexRepo.GetPokemonDexByPokemonDexName(&pokemonDex, pokemonDexName)
	if handle != nil {
		fmt.Println("Error")
	}
	return pokemonDex, handle
}
