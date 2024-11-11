package models

type RegistrationContextContainer struct {
	UeContext           UeContext          `json:"ueContext"`
	AnN2ApId            int                `json:"anN2ApId"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnType              string             `json:"anType"`
	InitialAmfName      string             `json:"initialAmfName"`
	UserLocation        UserLocation       `json:"userLocation"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
}
