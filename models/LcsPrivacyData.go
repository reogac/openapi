package models

type LcsPrivacyData struct {
	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`
	Lpi                 *Lpi                `json:"lpi,omitempty"`
	UnrelatedClass      *UnrelatedClass     `json:"unrelatedClass,omitempty"`
}
