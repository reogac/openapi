package models
type PointUncertaintyEllipse struct {
	 Shape	SupportedGADShapes	`json:"shape"`
	 Confidence	int	`json:"confidence"`
	 Point	GeographicalCoordinates	`json:"point"`
	 UncertaintyEllipse	UncertaintyEllipse	`json:"uncertaintyEllipse"`
}
