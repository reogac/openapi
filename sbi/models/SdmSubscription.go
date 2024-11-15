/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
}
