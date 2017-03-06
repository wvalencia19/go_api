package models

const (
	// CoordinatesPrecision The Precision of Float to be Used on Coordinates
	CoordinatesPrecision = 6
	// CoordinatesBitSize The System Architecture to be Used on Coordinates
	CoordinatesBitSize = 64
)

// Coordinates Entity Including Latitude and Longitude Data
type Coordinates struct {
	Latitude  float64
	Longitude float64
}
