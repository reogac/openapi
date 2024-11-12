package models

type SdmSubscription struct {
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	Dnn                        string                       `json:"dnn,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
}
