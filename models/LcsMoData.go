package models
type LcsMoData struct {
	 MoAssistanceDataTypes	*LcsBroadcastAssistanceTypesData	`json:"moAssistanceDataTypes,omitempty"`
	 AllowedServiceClasses	[]string	`json:"allowedServiceClasses"`
}
