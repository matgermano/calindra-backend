package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	baseEndpoint = "https://maps.googleapis.com/maps/api/geocode/json?address="
)

func GetCoordinatesFromAddress(address string) (*Geo, error) {
	uri := url.QueryEscape(address)
	endpoint := baseEndpoint + uri
	finalEndpoint := endpoint + *key

	x, err := http.Get(finalEndpoint)
	if err != nil {
		return nil, err
	}
	defer x.Body.Close()

	var geo Geo
	err = json.NewDecoder(x.Body).Decode(&geo)
	if err != nil {
		return nil, err
	}

	return &geo, nil
}
