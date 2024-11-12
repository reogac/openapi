package models

type SmfSelectionSubscriptionData struct {
	SubscribedSnssaiInfos map[string]SnssaiInfo `json:"subscribedSnssaiInfos,omitempty"`
	SharedSnssaiInfosId   string                `json:"sharedSnssaiInfosId,omitempty"`
	HssGroupId            string                `json:"hssGroupId,omitempty"`
	SupportedFeatures     string                `json:"supportedFeatures,omitempty"`
}
