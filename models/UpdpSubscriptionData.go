package models

type UpdpSubscriptionData struct {
	UpdpNotifyCallbackUri    string `json:"updpNotifyCallbackUri"`
	SupportedFeatures        string `json:"supportedFeatures,omitempty"`
	UpdpCallbackBinding      string `json:"updpCallbackBinding,omitempty"`
	UpdpNotifySubscriptionId string `json:"updpNotifySubscriptionId"`
}
