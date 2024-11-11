package models
type UeN1N2InfoSubscriptionCreateData struct {
	 N2InformationClass	string	`json:"n2InformationClass,omitempty"`
	 N2NotifyCallbackUri	string	`json:"n2NotifyCallbackUri,omitempty"`
	 N1MessageClass	string	`json:"n1MessageClass,omitempty"`
	 N1NotifyCallbackUri	string	`json:"n1NotifyCallbackUri,omitempty"`
	 NfId	string	`json:"nfId,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 OldGuami	*Guami	`json:"oldGuami,omitempty"`
}
