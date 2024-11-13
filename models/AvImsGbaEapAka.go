package models

type AvImsGbaEapAka struct {
	Autn   string    `json:"autn"`
	Ck     string    `json:"ck"`
	Ik     string    `json:"ik"`
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
}
