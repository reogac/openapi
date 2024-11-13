package models

type GeographicArea struct {
	PointUncertaintyEllipse  *PointUncertaintyEllipse  `json:"PointUncertaintyEllipse,omitempty"`
	Polygon                  *Polygon                  `json:"Polygon,omitempty"`
	PointAltitude            *PointAltitude            `json:"PointAltitude,omitempty"`
	PointAltitudeUncertainty *PointAltitudeUncertainty `json:"PointAltitudeUncertainty,omitempty"`
	EllipsoidArc             *EllipsoidArc             `json:"EllipsoidArc,omitempty"`
	Point                    *Point                    `json:"Point,omitempty"`
	PointUncertaintyCircle   *PointUncertaintyCircle   `json:"PointUncertaintyCircle,omitempty"`
}
