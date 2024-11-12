package models

type PointUncertaintyEllipse struct {
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	Confidence         int                     `json:"confidence"`
	Shape              SupportedGADShapes      `json:"shape"`
	Point              GeographicalCoordinates `json:"point"`
}
