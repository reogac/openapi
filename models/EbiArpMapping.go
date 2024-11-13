package models

type EbiArpMapping struct {
	Arp         Arp `json:"arp"`
	EpsBearerId int `json:"epsBearerId"`
}
