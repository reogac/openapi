package models

type SdmSubscription struct {
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	Dnn                        string                       `json:"dnn,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	Expires                    string                       `json:"expires,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
}
