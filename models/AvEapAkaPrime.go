package models

type AvEapAkaPrime struct {
	Xres    string `json:"xres"`
	Autn    string `json:"autn"`
	CkPrime string `json:"ckPrime"`
	IkPrime string `json:"ikPrime"`
	AvType  AvType `json:"avType"`
	Rand    string `json:"rand"`
}
