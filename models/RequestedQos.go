package models

type RequestedQos struct {
	GbrUl  string `json:"gbrUl,omitempty"`
	GbrDl  string `json:"gbrDl,omitempty"`
	FiveQi int    `json:"5qi"`
}
