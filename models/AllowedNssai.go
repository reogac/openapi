package models

type AllowedNssai struct {
	AllowedSnssaiList []AllowedSnssai `json:"allowedSnssaiList"`
	AccessType        string          `json:"accessType"`
}
