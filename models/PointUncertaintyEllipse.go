package models

type PointUncertaintyEllipse struct {
	Shape              SupportedGADShapes      `json:"shape"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	Confidence         int                     `json:"confidence"`
	Point              GeographicalCoordinates `json:"point"`
}
