/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	Expires                    string                       `json:"expires,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
}
