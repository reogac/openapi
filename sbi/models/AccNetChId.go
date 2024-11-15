/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccNetChId struct {
	SessionChScope   *bool    `json:"sessionChScope,omitempty"`
	AccNetChaIdValue *int     `json:"accNetChaIdValue,omitempty"`
	AccNetChargId    string   `json:"accNetChargId,omitempty"`
	RefPccRuleIds    []string `json:"refPccRuleIds,omitempty"`
}
