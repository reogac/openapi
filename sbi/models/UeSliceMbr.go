/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:40 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSliceMbr struct {
	SliceMbr         map[string]SliceMbr `json:"sliceMbr"`
	ServingSnssai    Snssai              `json:"servingSnssai"`
	MappedHomeSnssai *Snssai             `json:"mappedHomeSnssai,omitempty"`
}
