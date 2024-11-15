/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:28 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NumberAverage struct {
	Variance float64  `json:"variance"`
	Skewness *float64 `json:"skewness,omitempty"`
	Number   float64  `json:"number"`
}
