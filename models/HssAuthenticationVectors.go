package models

type HssAuthenticationVectors struct {
	AvImsGbaEapAka []AvImsGbaEapAka `json:"AvImsGbaEapAka,omitempty"`
	AvEapAkaPrime  []AvEapAkaPrime  `json:"AvEapAkaPrime,omitempty"`
	AvEpsAka       []AvEpsAka       `json:"AvEpsAka,omitempty"`
}
