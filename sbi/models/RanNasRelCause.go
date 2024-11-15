/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanNasRelCause struct {
	FiveGSmCause *int       `json:"5gSmCause,omitempty"`
	EpsCause     string     `json:"epsCause,omitempty"`
	NgApCause    *NgApCause `json:"ngApCause,omitempty"`
	FiveGMmCause *int       `json:"5gMmCause,omitempty"`
}
