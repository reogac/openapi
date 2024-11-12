package models

type PointUncertaintyEllipse struct {
	Shape              SupportedGADShapes      `json:"shape"`
	Point              GeographicalCoordinates `json:"point"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	Confidence         int                     `json:"confidence"`
}
