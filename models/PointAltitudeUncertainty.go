package models

type PointAltitudeUncertainty struct {
	UncertaintyAltitude float64                 `json:"uncertaintyAltitude"`
	Confidence          int                     `json:"confidence"`
	VConfidence         *int                    `json:"vConfidence,omitempty"`
	Shape               string                  `json:"shape"`
	Point               GeographicalCoordinates `json:"point"`
	Altitude            float64                 `json:"altitude"`
	UncertaintyEllipse  UncertaintyEllipse      `json:"uncertaintyEllipse"`
}
