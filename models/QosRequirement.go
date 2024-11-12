package models
type QosRequirement struct {
	 ResType	QosResourceType	`json:"resType,omitempty"`
	 Pdb	*int	`json:"pdb,omitempty"`
	 Per	string	`json:"per,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
	 GfbrUl	string	`json:"gfbrUl,omitempty"`
	 GfbrDl	string	`json:"gfbrDl,omitempty"`
}
