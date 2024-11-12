package models

type SdmSubscription struct {
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	Expires                    string                       `json:"expires,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
}
