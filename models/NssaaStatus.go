package models

type NssaaStatus struct {
	Snssai Snssai `json:"snssai"`
	Status string `json:"status"`
}
