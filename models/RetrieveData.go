package models

type RetrieveData struct {
	SmallDataRateStatusReq *bool                 `json:"smallDataRateStatusReq,omitempty"`
	PduSessionContextType  PduSessionContextType `json:"pduSessionContextType,omitempty"`
}
