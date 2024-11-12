package models

type RetrieveData struct {
	PduSessionContextType  PduSessionContextType `json:"pduSessionContextType,omitempty"`
	SmallDataRateStatusReq *bool                 `json:"smallDataRateStatusReq,omitempty"`
}
