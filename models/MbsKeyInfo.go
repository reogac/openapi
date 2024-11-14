/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsKeyInfo struct {
	MtkId       string `json:"mtkId,omitempty"`
	Mtk         string `json:"mtk,omitempty"`
	KeyDomainId string `json:"keyDomainId"`
	MskId       string `json:"mskId"`
	Msk         string `json:"msk,omitempty"`
	MskLifetime string `json:"mskLifetime,omitempty"`
}
