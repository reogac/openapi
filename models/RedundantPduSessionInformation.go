package models

type RedundantPduSessionInformation struct {
	Rsn              Rsn  `json:"rsn"`
	PduSessionPairId *int `json:"pduSessionPairId,omitempty"`
}
