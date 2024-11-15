/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	Pei                   string                 `json:"pei,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
}
