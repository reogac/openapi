package models

type NonUeN2InfoSubscriptionCreateData struct {
	NfId                string             `json:"nfId,omitempty"`
	SupportedFeatures   string             `json:"supportedFeatures,omitempty"`
	GlobalRanNodeList   []GlobalRanNodeId  `json:"globalRanNodeList,omitempty"`
	AnTypeList          []string           `json:"anTypeList,omitempty"`
	N2InformationClass  N2InformationClass `json:"n2InformationClass"`
	N2NotifyCallbackUri string             `json:"n2NotifyCallbackUri"`
}
