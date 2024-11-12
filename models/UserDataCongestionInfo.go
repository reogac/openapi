package models

type UserDataCongestionInfo struct {
	Snssai         *Snssai         `json:"snssai,omitempty"`
	NetworkArea    NetworkAreaInfo `json:"networkArea"`
	CongestionInfo CongestionInfo  `json:"congestionInfo"`
}
