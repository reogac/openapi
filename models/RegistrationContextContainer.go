package models

type RegistrationContextContainer struct {
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	AnType              AccessType         `json:"anType"`
	AnN2ApId            int                `json:"anN2ApId"`
	UserLocation        UserLocation       `json:"userLocation"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	InitialAmfName      string             `json:"initialAmfName"`
}
