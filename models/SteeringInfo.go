/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SteeringInfo struct {
	PlmnId         PlmnId   `json:"plmnId"`
	AccessTechList []string `json:"accessTechList,omitempty"`
}
