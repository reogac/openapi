package models

type AmfEventSubscriptionAddInfo struct {
	EventSyncInd      *bool    `json:"eventSyncInd,omitempty"`
	NfConsumerInfo    []string `json:"nfConsumerInfo,omitempty"`
	BindingInfo       []string `json:"bindingInfo,omitempty"`
	SubscribingNfType string   `json:"subscribingNfType,omitempty"`
}
