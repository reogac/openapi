/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	Dnn                        string                       `json:"dnn,omitempty"`
}
