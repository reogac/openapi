package models

type PointAltitude struct {
	Altitude float64                 `json:"altitude"`
	Shape    SupportedGADShapes      `json:"shape"`
	Point    GeographicalCoordinates `json:"point"`
}
