package models
type PointAltitude struct {
	 Point	GeographicalCoordinates	`json:"point"`
	 Altitude	float64	`json:"altitude"`
	 Shape	SupportedGADShapes	`json:"shape"`
}
