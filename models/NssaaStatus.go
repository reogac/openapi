package models

type NssaaStatus struct {
	Status AuthStatus `json:"status"`
	Snssai Snssai     `json:"snssai"`
}
