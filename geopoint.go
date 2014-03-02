package geopoint

import (
	"math"
)

const (
	mileInKilometres        = 1.60934         // kilometres -> miles conversion.
	degToRad                = math.Pi / 180.0 // degrees -> radians conversion.
	earthRadiusInKilometres = 6371
)

// Kilometres are units of distance.
type Kilometres float64

// Miles converts kilometres to miles.
func (self Kilometres) Miles() Miles {
	return Miles(self * mileInKilometres)
}

// Miles are units of distance.
type Miles float64

// Kilometres converts miles to kilometres.
func (self Miles) Kilometres() Kilometres {
	return Kilometres(self * mileInKilometres)
}

// Formula to calculate distance between two GeoPoints.
type Formula func(one, two *GeoPoint) Kilometres

// Degrees are units of angular measure.
type Degrees float64

// Conversion between Degrees and Radians.
func (self Degrees) Radians() Radians {
	return Radians(self * degToRad)
}

// Radians are units of angular measure.
type Radians float64

// Conversion between Radians and Degrees.
func (self Radians) Degrees() Degrees {
	return Degrees(self / degToRad)
}

// Struct representing GPS coordinates.
type GeoPoint struct {
	Latitude, Longitude Degrees
}

// Constructor for a GeoPoint.
func NewGeoPoint(latitude, longitude Degrees) *GeoPoint {
	return &GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// Distance to another GeoPoint.
func (g *GeoPoint) DistanceTo(another *GeoPoint, f Formula) Kilometres {
	return f(g, another)
}

// Distance between two GeoPoints using Haversine formula.
func Haversine(one, two *GeoPoint) Kilometres {
	lat1 := one.Latitude.Radians()
	lng1 := one.Longitude.Radians()
	lat2 := two.Latitude.Radians()
	lng2 := two.Longitude.Radians()

	deltaLng := float64(lng2 - lng1)
	deltaLat := float64(lat2 - lat1)
	a := math.Pow((math.Sin(deltaLat/2)), 2.0) + math.Cos(float64(lat1))*
		math.Cos(float64(lat2))*math.Pow(math.Sin(deltaLng/2), 2.0)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	return Kilometres(earthRadiusInKilometres * c)
}
