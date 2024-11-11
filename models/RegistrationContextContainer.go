package models

type RegistrationContextContainer struct {
	UserLocation        UserLocation       `json:"userLocation"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	AnType              string             `json:"anType"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
}
