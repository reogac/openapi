package models
type RequestedQos struct {
	 FiveQi	int	`json:"5qi"`
	 GbrUl	string	`json:"gbrUl,omitempty"`
	 GbrDl	string	`json:"gbrDl,omitempty"`
}
