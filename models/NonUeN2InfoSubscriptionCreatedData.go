package models
type NonUeN2InfoSubscriptionCreatedData struct {
	 N2NotifySubscriptionId	string	`json:"n2NotifySubscriptionId"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 N2InformationClass	N2InformationClass	`json:"n2InformationClass,omitempty"`
}
