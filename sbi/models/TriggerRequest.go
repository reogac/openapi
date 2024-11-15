/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TriggerRequest struct {
	Supi        string        `json:"supi"`
	FailedPcscf *PcscfAddress `json:"failedPcscf,omitempty"`
}
