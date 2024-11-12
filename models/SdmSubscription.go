package models

type SdmSubscription struct {
	CallbackReference          string                       `json:"callbackReference"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	Expires                    string                       `json:"expires,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
}
