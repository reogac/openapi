package models
type AvEpsAka struct {
	 Rand	string	`json:"rand"`
	 Xres	string	`json:"xres"`
	 Autn	string	`json:"autn"`
	 Kasme	string	`json:"kasme"`
	 AvType	HssAvType	`json:"avType"`
}
