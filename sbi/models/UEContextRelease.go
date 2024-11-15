/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UEContextRelease struct {
	Supi                string    `json:"supi,omitempty"`
	UnauthenticatedSupi *bool     `json:"unauthenticatedSupi,omitempty"`
	NgapCause           NgApCause `json:"ngapCause"`
}
