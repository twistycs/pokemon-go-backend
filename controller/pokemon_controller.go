package controller

import (
	"net/http"

	"github.com/twistycs/pokemon-go-backend/imp"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	pokemonService imp.PokemonService
}

func PokemonControllerInit(pokemonService imp.PokemonService) *PokemonController {
	return &PokemonController{
		pokemonService: pokemonService,
	}
}
func (u *PokemonController) GetAllPokemonController(c *gin.Context) {
	resp, err := u.pokemonService.GetAllPokemon()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *PokemonController) GetPokemonByNameController(c *gin.Context) {
	pokemonName := c.Params.ByName("pokemonName")
	resp, err := u.pokemonService.GetPokemonByPokemonName(pokemonName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
