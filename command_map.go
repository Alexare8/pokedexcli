package main

import (
	"fmt"
)

func commandMap(conf *config) error {
	response, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.prevLocationsURL = response.Previous
	conf.nextLocationsURL = response.Next

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(conf *config) error {
	if conf.prevLocationsURL == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	response, err := conf.pokeapiClient.ListLocations(conf.prevLocationsURL)
	if err != nil {
		return err
	}

	conf.prevLocationsURL = response.Previous
	conf.nextLocationsURL = response.Next

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}
