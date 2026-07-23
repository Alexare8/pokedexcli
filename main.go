package main

import (
	"time"

	"github.com/Alexare8/pokedexcli/internal/pokeapi"
)

func main() {
	conf := config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	startRepl(&conf)
}
