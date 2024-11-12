package models

type RedundantTransmissionExpReq struct {
	RedTOrderCriter string `json:"redTOrderCriter,omitempty"`
	Order           string `json:"order,omitempty"`
}
