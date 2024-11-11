package models

type QosRequirement struct {
	FiveQi  *int   `json:"5qi,omitempty"`
	GfbrUl  string `json:"gfbrUl,omitempty"`
	GfbrDl  string `json:"gfbrDl,omitempty"`
	ResType string `json:"resType,omitempty"`
	Pdb     *int   `json:"pdb,omitempty"`
	Per     string `json:"per,omitempty"`
}
