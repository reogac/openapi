package models

type AllowedSnssai struct {
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
	MappedHomeSnssai   *Snssai          `json:"mappedHomeSnssai,omitempty"`
	AllowedSnssai      Snssai           `json:"allowedSnssai"`
}
