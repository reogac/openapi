package models

type UncertaintyEllipse struct {
	OrientationMajor int     `json:"orientationMajor"`
	SemiMajor        float64 `json:"semiMajor"`
	SemiMinor        float64 `json:"semiMinor"`
}
