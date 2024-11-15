/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	InitialAmfName      string             `json:"initialAmfName"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AnType              AccessType         `json:"anType"`
	UserLocation        UserLocation       `json:"userLocation"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
}
