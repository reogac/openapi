package models

type SdmSubscription struct {
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
}
