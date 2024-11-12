package models

type Point struct {
	Point GeographicalCoordinates `json:"point"`
	Shape SupportedGADShapes      `json:"shape"`
}
