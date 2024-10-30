package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type httpClient struct {
	address string
}

func HttpClient(reqAddress string) httpClient {
	return httpClient{
		address: reqAddress,
	}
}

func (client *httpClient) GetShips() ([]Ship, error) {
	url := client.address + "/ships"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	shipsMock := make([]Ship, 0)

	shipsMock = append(shipsMock, Ship{
		Shipnr:       12313,
		Name:         "asd",
		Owner:        2,
		ShipType:     "sdasd",
		Currentvalue: "dasdsa",
		Image:        "asdasda",
		Year:         "qsadas",
		Crewmembers:  make([]CrewMember, 0),
	})
	return shipsMock, nil

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	var ships []Ship

	if err := json.Unmarshal(body, &ships); err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ships, nil
}

func (client *httpClient) SaveShips(ships []Ship) {

}
