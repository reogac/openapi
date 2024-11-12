package models

type SmsSubscriptionData struct {
	SharedSmsSubsDataId string `json:"sharedSmsSubsDataId,omitempty"`
	SupportedFeatures   string `json:"supportedFeatures,omitempty"`
	SmsSubscribed       *bool  `json:"smsSubscribed,omitempty"`
}
