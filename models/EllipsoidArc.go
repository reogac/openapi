package models

type EllipsoidArc struct {
	OffsetAngle       int                     `json:"offsetAngle"`
	IncludedAngle     int                     `json:"includedAngle"`
	Confidence        int                     `json:"confidence"`
	Point             GeographicalCoordinates `json:"point"`
	InnerRadius       int32                   `json:"innerRadius"`
	Shape             SupportedGADShapes      `json:"shape"`
	UncertaintyRadius float64                 `json:"uncertaintyRadius"`
}
