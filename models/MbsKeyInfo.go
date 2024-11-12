package models

type MbsKeyInfo struct {
	MtkId       string `json:"mtkId,omitempty"`
	Mtk         string `json:"mtk,omitempty"`
	KeyDomainId string `json:"keyDomainId"`
	MskId       string `json:"mskId"`
	Msk         string `json:"msk,omitempty"`
	MskLifetime string `json:"mskLifetime,omitempty"`
}
