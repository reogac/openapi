/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
}
