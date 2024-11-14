package models
type RoamingInfoUpdate struct {
	 Roaming	*bool	`json:"roaming,omitempty"`
	 ServingPlmn	PlmnId	`json:"servingPlmn"`
}
