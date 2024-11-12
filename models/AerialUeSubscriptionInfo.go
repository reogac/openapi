package models

type AerialUeSubscriptionInfo struct {
	ThreeGppUavId string             `json:"3gppUavId,omitempty"`
	AerialUeInd   AerialUeIndication `json:"aerialUeInd"`
}
