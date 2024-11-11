package models
type NonUeN2InfoSubscriptionCreateData struct {
	 GlobalRanNodeList	[]GlobalRanNodeId	`json:"globalRanNodeList,omitempty"`
	 AnTypeList	[]string	`json:"anTypeList,omitempty"`
	 N2InformationClass	string	`json:"n2InformationClass"`
	 N2NotifyCallbackUri	string	`json:"n2NotifyCallbackUri"`
	 NfId	string	`json:"nfId,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
