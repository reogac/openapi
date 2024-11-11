package models

type PointUncertaintyCircle struct {
	Shape       string                  `json:"shape"`
	Point       GeographicalCoordinates `json:"point"`
	Uncertainty float64                 `json:"uncertainty"`
}
