/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeographicArea struct {
	EllipsoidArc             *EllipsoidArc             `json:"EllipsoidArc,omitempty"`
	Point                    *Point                    `json:"Point,omitempty"`
	PointUncertaintyCircle   *PointUncertaintyCircle   `json:"PointUncertaintyCircle,omitempty"`
	PointUncertaintyEllipse  *PointUncertaintyEllipse  `json:"PointUncertaintyEllipse,omitempty"`
	Polygon                  *Polygon                  `json:"Polygon,omitempty"`
	PointAltitude            *PointAltitude            `json:"PointAltitude,omitempty"`
	PointAltitudeUncertainty *PointAltitudeUncertainty `json:"PointAltitudeUncertainty,omitempty"`
}
