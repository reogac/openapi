package models

type PointUncertaintyCircle struct {
	Point       GeographicalCoordinates `json:"point"`
	Uncertainty float64                 `json:"uncertainty"`
	Shape       SupportedGADShapes      `json:"shape"`
}
