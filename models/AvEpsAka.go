package models

type AvEpsAka struct {
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
	Autn   string    `json:"autn"`
	Kasme  string    `json:"kasme"`
}
