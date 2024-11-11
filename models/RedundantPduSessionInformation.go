package models

type RedundantPduSessionInformation struct {
	Rsn              string `json:"rsn"`
	PduSessionPairId *int   `json:"pduSessionPairId,omitempty"`
}
