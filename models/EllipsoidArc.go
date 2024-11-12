package models

type EllipsoidArc struct {
	InnerRadius       int32                   `json:"innerRadius"`
	UncertaintyRadius float64                 `json:"uncertaintyRadius"`
	Shape             SupportedGADShapes      `json:"shape"`
	OffsetAngle       int                     `json:"offsetAngle"`
	IncludedAngle     int                     `json:"includedAngle"`
	Confidence        int                     `json:"confidence"`
	Point             GeographicalCoordinates `json:"point"`
}
