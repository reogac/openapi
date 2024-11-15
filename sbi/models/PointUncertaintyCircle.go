/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointUncertaintyCircle struct {
	Uncertainty float64                 `json:"uncertainty"`
	Shape       SupportedGADShapes      `json:"shape"`
	Point       GeographicalCoordinates `json:"point"`
}
