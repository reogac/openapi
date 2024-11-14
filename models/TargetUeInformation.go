/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TargetUeInformation struct {
	AnyUe       *bool    `json:"anyUe,omitempty"`
	Supis       []string `json:"supis,omitempty"`
	Gpsis       []string `json:"gpsis,omitempty"`
	IntGroupIds []string `json:"intGroupIds,omitempty"`
}
