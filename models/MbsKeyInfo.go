package models

type MbsKeyInfo struct {
	MskId       string `json:"mskId"`
	Msk         string `json:"msk,omitempty"`
	MskLifetime string `json:"mskLifetime,omitempty"`
	MtkId       string `json:"mtkId,omitempty"`
	Mtk         string `json:"mtk,omitempty"`
	KeyDomainId string `json:"keyDomainId"`
}
