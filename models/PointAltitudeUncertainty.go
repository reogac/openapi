package models
type PointAltitudeUncertainty struct {
	 UncertaintyEllipse	UncertaintyEllipse	`json:"uncertaintyEllipse"`
	 UncertaintyAltitude	float64	`json:"uncertaintyAltitude"`
	 Confidence	int	`json:"confidence"`
	 VConfidence	*int	`json:"vConfidence,omitempty"`
	 Shape	SupportedGADShapes	`json:"shape"`
	 Point	GeographicalCoordinates	`json:"point"`
	 Altitude	float64	`json:"altitude"`
}
