package models

type UncertaintyEllipse struct {
	SemiMajor        float64 `json:"semiMajor"`
	SemiMinor        float64 `json:"semiMinor"`
	OrientationMajor int     `json:"orientationMajor"`
}
