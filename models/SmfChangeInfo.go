package models

type SmfChangeInfo struct {
	PduSessionIdList []int               `json:"pduSessionIdList"`
	SmfChangeInd     SmfChangeIndication `json:"smfChangeInd"`
}
