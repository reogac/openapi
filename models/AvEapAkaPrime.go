package models
type AvEapAkaPrime struct {
	 Rand	string	`json:"rand"`
	 Xres	string	`json:"xres"`
	 Autn	string	`json:"autn"`
	 CkPrime	string	`json:"ckPrime"`
	 IkPrime	string	`json:"ikPrime"`
	 AvType	AvType	`json:"avType"`
}
