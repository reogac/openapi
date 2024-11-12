package models

type NsiIdInfo struct {
	NsiIds []string `json:"nsiIds,omitempty"`
	Snssai Snssai   `json:"snssai"`
}
