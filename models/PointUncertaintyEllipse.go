package models

type PointUncertaintyEllipse struct {
	Shape              string                  `json:"shape"`
	Point              GeographicalCoordinates `json:"point"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	Confidence         int                     `json:"confidence"`
}
