package main

import (
	"fmt"
)

func commandMap(conf *config) error {
	response, err := conf.pokeapiClient.ListLocations(conf.Next)
	if err != nil {
		return err
	}

	conf.Previous = response.Previous
	conf.Next = response.Next

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(conf *config) error {
	if conf.Previous == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	response, err := conf.pokeapiClient.ListLocations(conf.Previous)
	if err != nil {
		return err
	}

	conf.Previous = response.Previous
	conf.Next = response.Next

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}
