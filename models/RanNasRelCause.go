package models

type RanNasRelCause struct {
	NgApCause    *NgApCause `json:"ngApCause,omitempty"`
	FiveGMmCause *int       `json:"5gMmCause,omitempty"`
	FiveGSmCause *int       `json:"5gSmCause,omitempty"`
	EpsCause     string     `json:"epsCause,omitempty"`
}