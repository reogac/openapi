/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:45 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
}
