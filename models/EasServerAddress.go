package models

type EasServerAddress struct {
	Port int    `json:"port"`
	Ip   IpAddr `json:"ip"`
}
