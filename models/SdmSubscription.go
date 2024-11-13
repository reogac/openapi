package models

type SdmSubscription struct {
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
}
