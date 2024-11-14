/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsMediaInfo struct {
	MbsMedType    MediaType `json:"mbsMedType,omitempty"`
	MaxReqMbsBwDl string    `json:"maxReqMbsBwDl,omitempty"`
	MinReqMbsBwDl string    `json:"minReqMbsBwDl,omitempty"`
	Codecs        []string  `json:"codecs,omitempty"`
}
