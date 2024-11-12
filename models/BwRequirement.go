package models

type BwRequirement struct {
	MirBwDl string `json:"mirBwDl,omitempty"`
	MirBwUl string `json:"mirBwUl,omitempty"`
	AppId   string `json:"appId"`
	MarBwDl string `json:"marBwDl,omitempty"`
	MarBwUl string `json:"marBwUl,omitempty"`
}
