package models

type UserDataCongestionInfo struct {
	CongestionInfo CongestionInfo  `json:"congestionInfo"`
	Snssai         *Snssai         `json:"snssai,omitempty"`
	NetworkArea    NetworkAreaInfo `json:"networkArea"`
}
