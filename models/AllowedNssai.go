package models

type AllowedNssai struct {
	AccessType        AccessType      `json:"accessType"`
	AllowedSnssaiList []AllowedSnssai `json:"allowedSnssaiList"`
}
