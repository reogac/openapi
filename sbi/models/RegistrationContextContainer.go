/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnType              AccessType         `json:"anType"`
}
