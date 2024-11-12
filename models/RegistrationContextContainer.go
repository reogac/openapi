package models

type RegistrationContextContainer struct {
	UeContext           UeContext          `json:"ueContext"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	InitialAmfName      string             `json:"initialAmfName"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	AnType              AccessType         `json:"anType"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
}
