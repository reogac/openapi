/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	AnType              AccessType         `json:"anType"`
	AnN2ApId            int                `json:"anN2ApId"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	InitialAmfName      string             `json:"initialAmfName"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
}
