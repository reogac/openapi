package models
type SdmSubscription struct {
	 Expires	string	`json:"expires,omitempty"`
	 AmfServiceName	ServiceName	`json:"amfServiceName,omitempty"`
	 MonitoredResourceUris	[]string	`json:"monitoredResourceUris"`
	 UniqueSubscription	*bool	`json:"uniqueSubscription,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 SingleNssai	*Snssai	`json:"singleNssai,omitempty"`
	 Report	*ImmediateReport	`json:"report,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 ImmediateReport	*bool	`json:"immediateReport,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 NfChangeFilter	*bool	`json:"nfChangeFilter,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 CallbackReference	string	`json:"callbackReference"`
	 SubscriptionId	string	`json:"subscriptionId,omitempty"`
	 PlmnId	*PlmnId	`json:"plmnId,omitempty"`
	 UeConSmfDataSubFilter	*UeContextInSmfDataSubFilter	`json:"ueConSmfDataSubFilter,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 ImplicitUnsubscribe	*bool	`json:"implicitUnsubscribe,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
}
