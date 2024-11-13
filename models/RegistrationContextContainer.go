package models

type RegistrationContextContainer struct {
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnType              AccessType         `json:"anType"`
	AnN2ApId            int                `json:"anN2ApId"`
	InitialAmfName      string             `json:"initialAmfName"`
	UeContext           UeContext          `json:"ueContext"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
}
