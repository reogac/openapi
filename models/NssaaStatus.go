package models

type NssaaStatus struct {
	Status string `json:"status"`
	Snssai Snssai `json:"snssai"`
}
