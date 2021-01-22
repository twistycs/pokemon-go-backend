package repositories

import (
	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
	"gorm.io/gorm"
)

type pokemonRepository struct {
	connect *gorm.DB
}

func PokemonRepositories(connect *gorm.DB) imp.PokemonRepository {
	return &pokemonRepository{connect}
}

func (repo *pokemonRepository) GetAllPokemon(u *[]models.Pokemon) (err error) {
	if err = repo.connect.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pokemonRepository) GetPokemonByPokemonName(u *models.Pokemon, pokemonName string) (err error) {
	if err = repo.connect.Where("name = ?", pokemonName).Find(u).Error; err != nil {
		return err
	}
	return nil
}
