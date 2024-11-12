package models

type Atom struct {
	Negative *bool  `json:"negative,omitempty"`
	Attr     string `json:"attr"`
}
