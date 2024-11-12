package models

type PointUncertaintyEllipse struct {
	Confidence         int                     `json:"confidence"`
	Point              GeographicalCoordinates `json:"point"`
	Shape              SupportedGADShapes      `json:"shape"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
}
