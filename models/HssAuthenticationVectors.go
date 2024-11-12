package models

type HssAuthenticationVectors struct {
	AvEpsAka       []AvEpsAka       `json:"AvEpsAka,omitempty"`
	AvImsGbaEapAka []AvImsGbaEapAka `json:"AvImsGbaEapAka,omitempty"`
	AvEapAkaPrime  []AvEapAkaPrime  `json:"AvEapAkaPrime,omitempty"`
}
