package models

type SdmSubscription struct {
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
}
