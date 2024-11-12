package models

type QosRequirement struct {
	GfbrUl  string `json:"gfbrUl,omitempty"`
	GfbrDl  string `json:"gfbrDl,omitempty"`
	ResType string `json:"resType,omitempty"`
	Pdb     *int   `json:"pdb,omitempty"`
	Per     string `json:"per,omitempty"`
	FiveQi  *int   `json:"5qi,omitempty"`
}
