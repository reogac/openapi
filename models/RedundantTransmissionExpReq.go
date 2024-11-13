package models

type RedundantTransmissionExpReq struct {
	Order           MatchingDirection            `json:"order,omitempty"`
	RedTOrderCriter RedTransExpOrderingCriterion `json:"redTOrderCriter,omitempty"`
}
