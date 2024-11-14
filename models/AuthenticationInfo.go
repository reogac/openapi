/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
}
