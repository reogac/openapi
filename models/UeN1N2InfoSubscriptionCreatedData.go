package models

type UeN1N2InfoSubscriptionCreatedData struct {
	SupportedFeatures        string `json:"supportedFeatures,omitempty"`
	N1n2NotifySubscriptionId string `json:"n1n2NotifySubscriptionId"`
}
