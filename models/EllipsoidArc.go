package models

type EllipsoidArc struct {
	InnerRadius       int32                   `json:"innerRadius"`
	UncertaintyRadius float64                 `json:"uncertaintyRadius"`
	OffsetAngle       int                     `json:"offsetAngle"`
	IncludedAngle     int                     `json:"includedAngle"`
	Shape             string                  `json:"shape"`
	Confidence        int                     `json:"confidence"`
	Point             GeographicalCoordinates `json:"point"`
}