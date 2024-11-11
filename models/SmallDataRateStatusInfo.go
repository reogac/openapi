package models

type SmallDataRateStatusInfo struct {
	Snssai              Snssai              `json:"Snssai"`
	Dnn                 string              `json:"Dnn"`
	SmallDataRateStatus SmallDataRateStatus `json:"SmallDataRateStatus"`
}
