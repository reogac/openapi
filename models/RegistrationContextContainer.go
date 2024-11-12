package models

type RegistrationContextContainer struct {
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnType              AccessType         `json:"anType"`
	InitialAmfName      string             `json:"initialAmfName"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	UserLocation        UserLocation       `json:"userLocation"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
}
