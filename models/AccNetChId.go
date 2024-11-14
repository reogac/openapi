/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccNetChId struct {
	AccNetChaIdValue *int     `json:"accNetChaIdValue,omitempty"`
	AccNetChargId    string   `json:"accNetChargId,omitempty"`
	RefPccRuleIds    []string `json:"refPccRuleIds,omitempty"`
	SessionChScope   *bool    `json:"sessionChScope,omitempty"`
}
