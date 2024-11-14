package models
type AdditionalSnssaiData struct {
	 RequiredAuthnAuthz	*bool	`json:"requiredAuthnAuthz,omitempty"`
	 SubscribedUeSliceMbr	*SliceMbr	`json:"subscribedUeSliceMbr,omitempty"`
	 SubscribedNsSrgList	[]string	`json:"subscribedNsSrgList,omitempty"`
}
