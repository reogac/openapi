package models
type UpdpSubscriptionData struct {
	 UpdpCallbackBinding	string	`json:"updpCallbackBinding,omitempty"`
	 UpdpNotifySubscriptionId	string	`json:"updpNotifySubscriptionId"`
	 UpdpNotifyCallbackUri	string	`json:"updpNotifyCallbackUri"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
