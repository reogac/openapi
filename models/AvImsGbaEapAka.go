package models

type AvImsGbaEapAka struct {
	Ck     string    `json:"ck"`
	Ik     string    `json:"ik"`
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
	Autn   string    `json:"autn"`
}
