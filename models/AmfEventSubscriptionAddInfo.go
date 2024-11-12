package models

type AmfEventSubscriptionAddInfo struct {
	NfConsumerInfo    []string                            `json:"nfConsumerInfo,omitempty"`
	AoiStateList      map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	BindingInfo       []string                            `json:"bindingInfo,omitempty"`
	SubscribingNfType string                              `json:"subscribingNfType,omitempty"`
	EventSyncInd      *bool                               `json:"eventSyncInd,omitempty"`
}
