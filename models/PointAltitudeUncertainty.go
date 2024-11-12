package models

type PointAltitudeUncertainty struct {
	Shape               SupportedGADShapes      `json:"shape"`
	Confidence          int                     `json:"confidence"`
	VConfidence         *int                    `json:"vConfidence,omitempty"`
	Point               GeographicalCoordinates `json:"point"`
	Altitude            float64                 `json:"altitude"`
	UncertaintyEllipse  UncertaintyEllipse      `json:"uncertaintyEllipse"`
	UncertaintyAltitude float64                 `json:"uncertaintyAltitude"`
}
