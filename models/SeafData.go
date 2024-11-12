package models

type SeafData struct {
	Ncc                  *int   `json:"ncc,omitempty"`
	KeyAmfChangeInd      *bool  `json:"keyAmfChangeInd,omitempty"`
	KeyAmfHDerivationInd *bool  `json:"keyAmfHDerivationInd,omitempty"`
	NgKsi                NgKsi  `json:"ngKsi"`
	KeyAmf               KeyAmf `json:"keyAmf"`
	Nh                   string `json:"nh,omitempty"`
}
