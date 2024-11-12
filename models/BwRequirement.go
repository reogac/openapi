package models

type BwRequirement struct {
	MarBwDl string `json:"marBwDl,omitempty"`
	MarBwUl string `json:"marBwUl,omitempty"`
	MirBwDl string `json:"mirBwDl,omitempty"`
	MirBwUl string `json:"mirBwUl,omitempty"`
	AppId   string `json:"appId"`
}
