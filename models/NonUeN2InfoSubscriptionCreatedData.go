package models

type NonUeN2InfoSubscriptionCreatedData struct {
	N2InformationClass     string `json:"n2InformationClass,omitempty"`
	N2NotifySubscriptionId string `json:"n2NotifySubscriptionId"`
	SupportedFeatures      string `json:"supportedFeatures,omitempty"`
}
