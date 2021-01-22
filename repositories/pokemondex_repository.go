package repositories

import (
	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pokemonDexRepository struct {
	connect *gorm.DB
}

func PokemonDexRepositories(connect *gorm.DB) imp.PokemonDexRepository {
	return &pokemonDexRepository{connect}
}

func (repo *pokemonDexRepository) GetAllPokemonDex(u *[]models.PokemonDex) (err error) {
	if err = repo.connect.Preload(clause.Associations).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pokemonDexRepository) GetPokemonDexByPokemonDexName(u *models.PokemonDex, pokemonDexName string) (err error) {
	if err = repo.connect.Where("name = ?", pokemonDexName).Find(u).Error; err != nil {
		return err
	}
	return nil
}
