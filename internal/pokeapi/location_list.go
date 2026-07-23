package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, exists := c.cache.Get(url)
	if !exists {
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		response, err := c.httpClient.Do(request)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer response.Body.Close()

		newData, err := io.ReadAll(response.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.cache.Add(url, newData)
		data = newData
	}

	response := RespShallowLocations{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return response, nil
}
