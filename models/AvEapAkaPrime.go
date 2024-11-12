package models

type AvEapAkaPrime struct {
	Autn    string `json:"autn"`
	CkPrime string `json:"ckPrime"`
	IkPrime string `json:"ikPrime"`
	AvType  AvType `json:"avType"`
	Rand    string `json:"rand"`
	Xres    string `json:"xres"`
}
