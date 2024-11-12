package models

type RegistrationContextContainer struct {
	UeContext           UeContext          `json:"ueContext"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	AnType              AccessType         `json:"anType"`
	InitialAmfName      string             `json:"initialAmfName"`
	UserLocation        UserLocation       `json:"userLocation"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
}
