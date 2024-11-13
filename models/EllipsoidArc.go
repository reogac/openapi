package models

type EllipsoidArc struct {
	UncertaintyRadius float64                 `json:"uncertaintyRadius"`
	OffsetAngle       int                     `json:"offsetAngle"`
	Shape             SupportedGADShapes      `json:"shape"`
	IncludedAngle     int                     `json:"includedAngle"`
	Confidence        int                     `json:"confidence"`
	Point             GeographicalCoordinates `json:"point"`
	InnerRadius       int32                   `json:"innerRadius"`
}
