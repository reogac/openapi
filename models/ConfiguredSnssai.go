package models

type ConfiguredSnssai struct {
	MappedHomeSnssai *Snssai `json:"mappedHomeSnssai,omitempty"`
	ConfiguredSnssai Snssai  `json:"configuredSnssai"`
}
