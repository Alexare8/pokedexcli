package main

import (
	"time"

	"github.com/Alexare8/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	conf := &config{
		pokeapiClient: client,
	}
	startRepl(conf)
}
