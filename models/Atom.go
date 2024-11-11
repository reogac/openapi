package models

type Atom struct {
	Attr     string `json:"attr"`
	Negative *bool  `json:"negative,omitempty"`
}
