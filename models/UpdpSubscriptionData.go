package models

type UpdpSubscriptionData struct {
	SupportedFeatures        string `json:"supportedFeatures,omitempty"`
	UpdpCallbackBinding      string `json:"updpCallbackBinding,omitempty"`
	UpdpNotifySubscriptionId string `json:"updpNotifySubscriptionId"`
	UpdpNotifyCallbackUri    string `json:"updpNotifyCallbackUri"`
}
