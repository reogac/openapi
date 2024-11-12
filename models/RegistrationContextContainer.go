package models

type RegistrationContextContainer struct {
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnType              AccessType         `json:"anType"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	InitialAmfName      string             `json:"initialAmfName"`
	UserLocation        UserLocation       `json:"userLocation"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
}
