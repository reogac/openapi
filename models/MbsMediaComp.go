package models

type MbsMediaComp struct {
	MbsMediaInfo  *MbsMediaInfo `json:"mbsMediaInfo,omitempty"`
	QosRef        string        `json:"qosRef,omitempty"`
	MbsQoSReq     *MbsQoSReq    `json:"mbsQoSReq,omitempty"`
	MbsMedCompNum int           `json:"mbsMedCompNum"`
	MbsFlowDescs  []string      `json:"mbsFlowDescs,omitempty"`
	MbsSdfResPrio string        `json:"mbsSdfResPrio,omitempty"`
}