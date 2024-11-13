package models

type RedundantTransmissionExpReq struct {
	RedTOrderCriter RedTransExpOrderingCriterion `json:"redTOrderCriter,omitempty"`
	Order           MatchingDirection            `json:"order,omitempty"`
}
