package models

type PointUncertaintyEllipse struct {
	Point              GeographicalCoordinates `json:"point"`
	Shape              string                  `json:"shape"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	Confidence         int                     `json:"confidence"`
}
