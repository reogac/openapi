/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosRequirement struct {
	FiveQi  *int            `json:"5qi,omitempty"`
	GfbrUl  string          `json:"gfbrUl,omitempty"`
	GfbrDl  string          `json:"gfbrDl,omitempty"`
	ResType QosResourceType `json:"resType,omitempty"`
	Pdb     *int            `json:"pdb,omitempty"`
	Per     string          `json:"per,omitempty"`
}
