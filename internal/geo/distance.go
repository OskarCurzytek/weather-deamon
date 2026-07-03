package geo

import (
	"math"

	"mod.com/m/internal/models"
)

func IsWithinRadius(c models.CityCoordinates, l *models.Lightning) bool {
	distance := haversineDistance(c, *l)
	return distance <= float64(c.Radius)
}

func haversineDistance(cords models.CityCoordinates, l models.Lightning) float64 {
	lat1 := degreesToRadians(cords.Lat)
	lon1 := degreesToRadians(cords.Lon)
	lat2 := degreesToRadians(l.Lat)
	lon2 := degreesToRadians(l.Lon)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * 6371
}

func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
