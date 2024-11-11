package models

type RedundantTransmissionExpReq struct {
	Order           string `json:"order,omitempty"`
	RedTOrderCriter string `json:"redTOrderCriter,omitempty"`
}
