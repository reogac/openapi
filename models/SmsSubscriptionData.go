package models
type SmsSubscriptionData struct {
	 SmsSubscribed	*bool	`json:"smsSubscribed,omitempty"`
	 SharedSmsSubsDataId	string	`json:"sharedSmsSubsDataId,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
