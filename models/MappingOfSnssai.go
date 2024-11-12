package models

type MappingOfSnssai struct {
	HomeSnssai    Snssai `json:"homeSnssai"`
	ServingSnssai Snssai `json:"servingSnssai"`
}
