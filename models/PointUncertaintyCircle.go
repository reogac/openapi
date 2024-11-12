package models

type PointUncertaintyCircle struct {
	Shape       SupportedGADShapes      `json:"shape"`
	Point       GeographicalCoordinates `json:"point"`
	Uncertainty float64                 `json:"uncertainty"`
}
