/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsServiceInfo struct {
	MbsMediaComps  map[string]MbsMediaComp `json:"mbsMediaComps"`
	MbsSdfResPrio  ReservPriority          `json:"mbsSdfResPrio,omitempty"`
	AfAppId        string                  `json:"afAppId,omitempty"`
	MbsSessionAmbr string                  `json:"mbsSessionAmbr,omitempty"`
}
