package models

type Point struct {
	Shape SupportedGADShapes      `json:"shape"`
	Point GeographicalCoordinates `json:"point"`
}
