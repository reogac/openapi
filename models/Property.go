/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Property struct {
	Name     string `json:"name"`
	Required *bool  `json:"required,omitempty"`
	Regex    string `json:"regex,omitempty"`
	Value    string `json:"value,omitempty"`
}
