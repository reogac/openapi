package models

type UpdpSubscriptionData struct {
	UpdpNotifySubscriptionId string `json:"updpNotifySubscriptionId"`
	UpdpNotifyCallbackUri    string `json:"updpNotifyCallbackUri"`
	SupportedFeatures        string `json:"supportedFeatures,omitempty"`
	UpdpCallbackBinding      string `json:"updpCallbackBinding,omitempty"`
}
