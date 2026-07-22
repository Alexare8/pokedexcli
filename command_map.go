package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(conf *config) error {
	url := conf.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := locationAreaResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	if response.Previous == nil {
		conf.Previous = ""
	} else {
		conf.Previous = *response.Previous
	}
	conf.Next = response.Next

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(conf *config) error {
	if conf.Previous == "" {
		fmt.Println("You're on the first page.")
		return nil
	}
	url := conf.Previous
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := locationAreaResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	if response.Previous == nil {
		conf.Previous = ""
	} else {
		conf.Previous = *response.Previous
	}
	conf.Next = response.Next

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}

type config struct {
	Next     string
	Previous string
}

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
