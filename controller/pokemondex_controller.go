package controller

import (
	"net/http"

	"github.com/twistycs/pokemon-go-backend/imp"

	"github.com/gin-gonic/gin"
)

type PokemonDexController struct {
	pokemonDexService imp.PokemonDexService
}

func PokemonDexControllerInit(pokemonDexService imp.PokemonDexService) *PokemonDexController {
	return &PokemonDexController{
		pokemonDexService: pokemonDexService,
	}
}
func (u *PokemonDexController) GetAllPokemonDexController(c *gin.Context) {
	resp, err := u.pokemonDexService.GetAllPokemonDex()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *PokemonDexController) GetPokemonDexByNameController(c *gin.Context) {
	pokemonDexName := c.Params.ByName("pokemonDexName")
	resp, err := u.pokemonDexService.GetPokemonDexByPokemonDexName(pokemonDexName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
