package main

import "math"

func EuclidianDistance(coordiante1, coordiante2 Location) (mi, km float64) {
	latitude1, longitude1 := coordiante1.ToRadians()
	latitude2, longitude2 := coordiante2.ToRadians()

	distance := haversineFormula(latitude1, latitude2, longitude1, longitude2)

	mi = distance * EarthRadiusMi
	km = distance * EarthRaidusKm
	return
}

func haversineFormula(latitude1, latitude2, longitude1, longitude2 float64) float64 {
	diffLat := latitude2 - latitude1
	diffLon := longitude2 - longitude1

	innerFormula := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(latitude1)*math.Cos(latitude2)*math.Pow(math.Sin(diffLon/2), 2)
	return 2 * math.Atan2(math.Sqrt(innerFormula), math.Sqrt(1-innerFormula))
}
