package models

type AdditionalSnssaiData struct {
	SubscribedUeSliceMbr *SliceMbr `json:"subscribedUeSliceMbr,omitempty"`
	SubscribedNsSrgList  []string  `json:"subscribedNsSrgList,omitempty"`
	RequiredAuthnAuthz   *bool     `json:"requiredAuthnAuthz,omitempty"`
}
