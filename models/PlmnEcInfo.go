package models

type PlmnEcInfo struct {
	EcRestrictionDataNb *bool                `json:"ecRestrictionDataNb,omitempty"`
	PlmnId              PlmnId               `json:"plmnId"`
	EcRestrictionDataWb *EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`
}
