package models

type PointAltitude struct {
	Shape    string                  `json:"shape"`
	Point    GeographicalCoordinates `json:"point"`
	Altitude float64                 `json:"altitude"`
}
