package models

type AbnormalBehaviour struct {
	Dnn          string                 `json:"dnn,omitempty"`
	Snssai       *Snssai                `json:"snssai,omitempty"`
	Ratio        *int                   `json:"ratio,omitempty"`
	Confidence   *int                   `json:"confidence,omitempty"`
	AddtMeasInfo *AdditionalMeasurement `json:"addtMeasInfo,omitempty"`
	Supis        []string               `json:"supis,omitempty"`
	Excep        Exception              `json:"excep"`
}
