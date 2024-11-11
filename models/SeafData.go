package models
type SeafData struct {
	 KeyAmf	KeyAmf	`json:"keyAmf"`
	 Nh	string	`json:"nh,omitempty"`
	 Ncc	*int	`json:"ncc,omitempty"`
	 KeyAmfChangeInd	*bool	`json:"keyAmfChangeInd,omitempty"`
	 KeyAmfHDerivationInd	*bool	`json:"keyAmfHDerivationInd,omitempty"`
	 NgKsi	NgKsi	`json:"ngKsi"`
}
