package models

type PointAltitude struct {
	Shape    SupportedGADShapes      `json:"shape"`
	Point    GeographicalCoordinates `json:"point"`
	Altitude float64                 `json:"altitude"`
}
