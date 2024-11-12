package models

type LcsMoData struct {
	AllowedServiceClasses []string                         `json:"allowedServiceClasses"`
	MoAssistanceDataTypes *LcsBroadcastAssistanceTypesData `json:"moAssistanceDataTypes,omitempty"`
}
