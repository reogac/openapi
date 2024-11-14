package models
type SdmSubscription struct {
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 AmfServiceName	ServiceName	`json:"amfServiceName,omitempty"`
	 SubscriptionId	string	`json:"subscriptionId,omitempty"`
	 ImmediateReport	*bool	`json:"immediateReport,omitempty"`
	 NfChangeFilter	*bool	`json:"nfChangeFilter,omitempty"`
	 UniqueSubscription	*bool	`json:"uniqueSubscription,omitempty"`
	 CallbackReference	string	`json:"callbackReference"`
	 Report	*ImmediateReport	`json:"report,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 ImplicitUnsubscribe	*bool	`json:"implicitUnsubscribe,omitempty"`
	 Expires	string	`json:"expires,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 MonitoredResourceUris	[]string	`json:"monitoredResourceUris"`
	 SingleNssai	*Snssai	`json:"singleNssai,omitempty"`
	 PlmnId	*PlmnId	`json:"plmnId,omitempty"`
	 UeConSmfDataSubFilter	*UeContextInSmfDataSubFilter	`json:"ueConSmfDataSubFilter,omitempty"`
}
