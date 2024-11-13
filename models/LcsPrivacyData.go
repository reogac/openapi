package models

type LcsPrivacyData struct {
	Lpi                 *Lpi                `json:"lpi,omitempty"`
	UnrelatedClass      *UnrelatedClass     `json:"unrelatedClass,omitempty"`
	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`
}
