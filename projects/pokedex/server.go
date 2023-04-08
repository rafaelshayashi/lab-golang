package main

import (
	"net/http"

	"github.com.br/rafaelshayashi/lab-golang/projects/pokedex/logger"
	"github.com.br/rafaelshayashi/lab-golang/projects/pokedex/pokemon"
	"github.com/labstack/echo/v4"
)

type ServerContext struct {
	echo.Context
}

func main() {

	logger.Info("starting the application")

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ServerContext{c}
			return next(cc)
		}
	})

	e.GET("/pokemons", listPokemons)
	e.POST("/pokemons", createPokemon)
	e.Logger.Fatal(e.Start(":1323"))
}

func listPokemons(c echo.Context) error {
	return c.JSON(http.StatusOK, pokemon.List())
}

func createPokemon(c echo.Context) error {
	return c.String(http.StatusOK, "should create a pokemon")
}
