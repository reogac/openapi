/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
}
