package models

type SmfChangeInfo struct {
	PduSessionIdList []int  `json:"pduSessionIdList"`
	SmfChangeInd     string `json:"smfChangeInd"`
}
