package models

type NonUeN2InfoSubscriptionCreatedData struct {
	SupportedFeatures      string             `json:"supportedFeatures,omitempty"`
	N2InformationClass     N2InformationClass `json:"n2InformationClass,omitempty"`
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
}
