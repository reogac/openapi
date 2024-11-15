/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnType              AccessType         `json:"anType"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnN2ApId            int                `json:"anN2ApId"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	UserLocation        UserLocation       `json:"userLocation"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
}
