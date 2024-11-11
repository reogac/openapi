package models

type AmfEventSubscriptionAddInfo struct {
	BindingInfo       []string `json:"bindingInfo,omitempty"`
	SubscribingNfType string   `json:"subscribingNfType,omitempty"`
	EventSyncInd      *bool    `json:"eventSyncInd,omitempty"`
	NfConsumerInfo    []string `json:"nfConsumerInfo,omitempty"`
}
