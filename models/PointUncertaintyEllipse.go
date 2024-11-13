package models

type PointUncertaintyEllipse struct {
	Point              GeographicalCoordinates `json:"point"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	Confidence         int                     `json:"confidence"`
	Shape              SupportedGADShapes      `json:"shape"`
}
