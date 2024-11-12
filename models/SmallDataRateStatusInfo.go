package models

type SmallDataRateStatusInfo struct {
	SmallDataRateStatus SmallDataRateStatus `json:"SmallDataRateStatus"`
	Snssai              Snssai              `json:"Snssai"`
	Dnn                 string              `json:"Dnn"`
}
