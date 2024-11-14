package models
type SmfSelectionSubscriptionData struct {
	 SharedSnssaiInfosId	string	`json:"sharedSnssaiInfosId,omitempty"`
	 HssGroupId	string	`json:"hssGroupId,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 SubscribedSnssaiInfos	map[string]SnssaiInfo	`json:"subscribedSnssaiInfos,omitempty"`
}
