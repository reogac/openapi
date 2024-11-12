package models

type LcsPrivacyData struct {
	UnrelatedClass      *UnrelatedClass     `json:"unrelatedClass,omitempty"`
	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`
	Lpi                 *Lpi                `json:"lpi,omitempty"`
}
