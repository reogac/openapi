package models

type RequestedQos struct {
	GbrDl  string `json:"gbrDl,omitempty"`
	FiveQi int    `json:"5qi"`
	GbrUl  string `json:"gbrUl,omitempty"`
}
