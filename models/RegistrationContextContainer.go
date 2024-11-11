package models

type RegistrationContextContainer struct {
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnType              string             `json:"anType"`
	UserLocation        UserLocation       `json:"userLocation"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
}
