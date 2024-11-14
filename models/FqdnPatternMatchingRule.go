/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FqdnPatternMatchingRule struct {
	Regex              string              `json:"regex,omitempty"`
	StringMatchingRule *StringMatchingRule `json:"stringMatchingRule,omitempty"`
}
