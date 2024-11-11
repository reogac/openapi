package models

type NssaiMapping struct {
	HSnssai      Snssai `json:"hSnssai"`
	MappedSnssai Snssai `json:"mappedSnssai"`
}
