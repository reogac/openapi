package models

type SmallDataRateStatusInfo struct {
	Dnn                 string              `json:"Dnn"`
	SmallDataRateStatus SmallDataRateStatus `json:"SmallDataRateStatus"`
	Snssai              Snssai              `json:"Snssai"`
}
