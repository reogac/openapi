package models
type EllipsoidArc struct {
	 OffsetAngle	int	`json:"offsetAngle"`
	 IncludedAngle	int	`json:"includedAngle"`
	 Confidence	int	`json:"confidence"`
	 Point	GeographicalCoordinates	`json:"point"`
	 Shape	SupportedGADShapes	`json:"shape"`
	 InnerRadius	int32	`json:"innerRadius"`
	 UncertaintyRadius	float64	`json:"uncertaintyRadius"`
}
