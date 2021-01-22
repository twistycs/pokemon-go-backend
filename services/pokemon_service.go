package services

import (
	"fmt"

	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
)

type pokemonService struct {
	pokemonRepo imp.PokemonRepository
}

func NewPokemonService(repo imp.PokemonRepository) imp.PokemonService {
	return &pokemonService{pokemonRepo: repo}
}

func (pokemonService *pokemonService) GetAllPokemon() (pokemons []models.Pokemon, err error) {
	var pokemon []models.Pokemon
	handle := pokemonService.pokemonRepo.GetAllPokemon(&pokemon)
	if handle != nil {
		fmt.Println("Error")
	}
	return pokemon, handle
}

func (pokemonService *pokemonService) GetPokemonByPokemonName(pokemonName string) (pokemons models.Pokemon, err error) {
	var pokemon models.Pokemon
	handle := pokemonService.pokemonRepo.GetPokemonByPokemonName(&pokemon, pokemonName)
	if handle != nil {
		fmt.Println("Error")
	}
	return pokemon, handle
}
