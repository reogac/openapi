/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SteeringInfo struct {
	PlmnId         PlmnId   `json:"plmnId"`
	AccessTechList []string `json:"accessTechList,omitempty"`
}
