package geopoint

import (
	"math"
)

const (
	EARTH_RADIUS_IN_KM = 6371
)

// Formula to calculate distance between two GeoPoints.
type Formula func(one, two *GeoPoint) float64

// Struct representing GPS coordinates.
type GeoPoint struct {
	Latitude, Longitude float64
}

// Constructor for a GeoPoint.
func NewGeoPoint(latitude, longitude float64) *GeoPoint {
	return &GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// Distance to another GeoPoint.
func (g *GeoPoint) DistanceTo(another *GeoPoint, f Formula) float64 {
	return f(g, another)
}

// Distance between two GeoPoints using Haversine formula.
func Haversine(one, two *GeoPoint) float64 {
	lat1 := toRadians(one.Latitude)
	lng1 := toRadians(one.Longitude)
	lat2 := toRadians(two.Latitude)
	lng2 := toRadians(two.Longitude)

	deltaLng := lng2 - lng1
	deltaLat := lat2 - lat1
	a := math.Pow((math.Sin(deltaLat/2)), 2.0) + math.Cos(lat1)*
		math.Cos(lat2)*math.Pow(math.Sin(deltaLng/2), 2.0)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	return EARTH_RADIUS_IN_KM * c
}

// Helper function to turn degrees into radians.
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
