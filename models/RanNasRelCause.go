package models

type RanNasRelCause struct {
	EpsCause     string     `json:"epsCause,omitempty"`
	NgApCause    *NgApCause `json:"ngApCause,omitempty"`
	FiveGMmCause *int       `json:"5gMmCause,omitempty"`
	FiveGSmCause *int       `json:"5gSmCause,omitempty"`
}
