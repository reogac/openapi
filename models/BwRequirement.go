package models
type BwRequirement struct {
	 MarBwUl	string	`json:"marBwUl,omitempty"`
	 MirBwDl	string	`json:"mirBwDl,omitempty"`
	 MirBwUl	string	`json:"mirBwUl,omitempty"`
	 AppId	string	`json:"appId"`
	 MarBwDl	string	`json:"marBwDl,omitempty"`
}
