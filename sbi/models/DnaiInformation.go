/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnaiInformation struct {
	Dnai                string `json:"dnai"`
	NoDnaiChangeInd     *bool  `json:"noDnaiChangeInd,omitempty"`
	NoLocalPsaChangeInd *bool  `json:"noLocalPsaChangeInd,omitempty"`
}
