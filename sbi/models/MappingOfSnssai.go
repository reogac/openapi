/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:56 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MappingOfSnssai struct {
	ServingSnssai Snssai `json:"servingSnssai"`
	HomeSnssai    Snssai `json:"homeSnssai"`
}
