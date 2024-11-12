package models

type EllipsoidArc struct {
	InnerRadius       int32                   `json:"innerRadius"`
	UncertaintyRadius float64                 `json:"uncertaintyRadius"`
	OffsetAngle       int                     `json:"offsetAngle"`
	IncludedAngle     int                     `json:"includedAngle"`
	Confidence        int                     `json:"confidence"`
	Shape             SupportedGADShapes      `json:"shape"`
	Point             GeographicalCoordinates `json:"point"`
}
