/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:08 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
}
