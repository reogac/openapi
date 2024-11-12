package models

type RedundantPduSessionInformation struct {
	PduSessionPairId *int   `json:"pduSessionPairId,omitempty"`
	Rsn              string `json:"rsn"`
}
