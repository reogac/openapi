package models

type RegistrationContextContainer struct {
	AnN2ApId            int                `json:"anN2ApId"`
	InitialAmfName      string             `json:"initialAmfName"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	AnType              AccessType         `json:"anType"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	UserLocation        UserLocation       `json:"userLocation"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
}
