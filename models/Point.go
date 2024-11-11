package models

type Point struct {
	Shape string                  `json:"shape"`
	Point GeographicalCoordinates `json:"point"`
}
