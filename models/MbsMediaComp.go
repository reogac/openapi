package models
type MbsMediaComp struct {
	 MbsSdfResPrio	ReservPriority	`json:"mbsSdfResPrio,omitempty"`
	 MbsMediaInfo	*MbsMediaInfo	`json:"mbsMediaInfo,omitempty"`
	 QosRef	string	`json:"qosRef,omitempty"`
	 MbsQoSReq	*MbsQoSReq	`json:"mbsQoSReq,omitempty"`
	 MbsMedCompNum	int	`json:"mbsMedCompNum"`
	 MbsFlowDescs	[]string	`json:"mbsFlowDescs,omitempty"`
}
