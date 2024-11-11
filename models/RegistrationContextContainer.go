package models

type RegistrationContextContainer struct {
	InitialAmfName      string             `json:"initialAmfName"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnType              string             `json:"anType"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	UserLocation        UserLocation       `json:"userLocation"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
}
