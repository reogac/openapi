/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	InitialAmfName      string             `json:"initialAmfName"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnType              AccessType         `json:"anType"`
	UserLocation        UserLocation       `json:"userLocation"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
}
