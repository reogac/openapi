package models

type PointAltitude struct {
	Altitude float64                 `json:"altitude"`
	Shape    string                  `json:"shape"`
	Point    GeographicalCoordinates `json:"point"`
}
