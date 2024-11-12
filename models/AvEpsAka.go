package models

type AvEpsAka struct {
	Kasme  string    `json:"kasme"`
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
	Autn   string    `json:"autn"`
}
