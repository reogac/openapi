package models

type UserDataCongestionInfo struct {
	NetworkArea    NetworkAreaInfo `json:"networkArea"`
	CongestionInfo CongestionInfo  `json:"congestionInfo"`
	Snssai         *Snssai         `json:"snssai,omitempty"`
}
