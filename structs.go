package main

import "math"

const (
	EarthRadiusMi = 3963 // EarthRadiusMi is the radius of the earth in miles.
	EarthRaidusKm = 6378 // EarthRaidusKm is the radius of the earth in kilometers.
)

type Geo struct {
	Results []Result `json:"results"`
}

type Result struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddres   string             `json:"formatted_address"`
	Geometry          Geometry           `json:"geometry"`
	PlaceID           string             `json:"place_id"`
	PlusCode          PlusCode           `json:"plus_code"`
	Types             []string           `json:"types"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Geometry struct {
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// ToRadians converts the given degree to radian
func (c Location) ToRadians() (latitude, longitude float64) {
	const constant = math.Pi / 180
	latitude = c.Latitude * constant
	longitude = c.Longitude * constant
	return
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Northwest Location `json:"northwest"`
	Southwest Location `json:"southwest"`
	Southeast Location `json:"southeast"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}
