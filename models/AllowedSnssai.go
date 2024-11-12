package models

type AllowedSnssai struct {
	MappedHomeSnssai   *Snssai          `json:"mappedHomeSnssai,omitempty"`
	AllowedSnssai      Snssai           `json:"allowedSnssai"`
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
}
