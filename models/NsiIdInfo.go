package models

type NsiIdInfo struct {
	Snssai Snssai   `json:"snssai"`
	NsiIds []string `json:"nsiIds,omitempty"`
}
