package models

type PointAltitudeUncertainty struct {
	UncertaintyAltitude float64                 `json:"uncertaintyAltitude"`
	Confidence          int                     `json:"confidence"`
	VConfidence         *int                    `json:"vConfidence,omitempty"`
	Point               GeographicalCoordinates `json:"point"`
	Altitude            float64                 `json:"altitude"`
	Shape               SupportedGADShapes      `json:"shape"`
	UncertaintyEllipse  UncertaintyEllipse      `json:"uncertaintyEllipse"`
}
