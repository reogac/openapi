package models

type RegistrationContextContainer struct {
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
	UserLocation        UserLocation       `json:"userLocation"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnType              AccessType         `json:"anType"`
}
