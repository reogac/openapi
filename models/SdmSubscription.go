package models

type SdmSubscription struct {
	Report                     *ImmediateReport             `json:"report,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
}
