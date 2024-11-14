package models
type AvEapAkaPrime struct {
	 IkPrime	string	`json:"ikPrime"`
	 AvType	AvType	`json:"avType"`
	 Rand	string	`json:"rand"`
	 Xres	string	`json:"xres"`
	 Autn	string	`json:"autn"`
	 CkPrime	string	`json:"ckPrime"`
}
