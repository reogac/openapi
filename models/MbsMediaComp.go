/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsMediaComp struct {
	QosRef        string         `json:"qosRef,omitempty"`
	MbsQoSReq     *MbsQoSReq     `json:"mbsQoSReq,omitempty"`
	MbsMedCompNum int            `json:"mbsMedCompNum"`
	MbsFlowDescs  []string       `json:"mbsFlowDescs,omitempty"`
	MbsSdfResPrio ReservPriority `json:"mbsSdfResPrio,omitempty"`
	MbsMediaInfo  *MbsMediaInfo  `json:"mbsMediaInfo,omitempty"`
}
