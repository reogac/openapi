package models
type MdtConfiguration struct {
	 CollectionPeriodRmmNr	string	`json:"collectionPeriodRmmNr,omitempty"`
	 InterFreqTargetList	[]InterFreqTargetInfo	`json:"interFreqTargetList,omitempty"`
	 EventThresholdRsrqNr	*int	`json:"eventThresholdRsrqNr,omitempty"`
	 LoggingDurationNr	string	`json:"loggingDurationNr,omitempty"`
	 SensorMeasurementList	[]string	`json:"sensorMeasurementList,omitempty"`
	 CollectionPeriodRmmLte	string	`json:"collectionPeriodRmmLte,omitempty"`
	 JobType	string	`json:"jobType"`
	 ReportAmount	string	`json:"reportAmount,omitempty"`
	 EventThresholdRsrpNr	*int	`json:"eventThresholdRsrpNr,omitempty"`
	 LoggingInterval	string	`json:"loggingInterval,omitempty"`
	 PositioningMethod	string	`json:"positioningMethod,omitempty"`
	 EventThresholdRsrq	*int	`json:"eventThresholdRsrq,omitempty"`
	 EventList	[]string	`json:"eventList,omitempty"`
	 AreaScope	*AreaScope	`json:"areaScope,omitempty"`
	 MeasurementNrList	[]string	`json:"measurementNrList,omitempty"`
	 LoggingDuration	string	`json:"loggingDuration,omitempty"`
	 ReportingTriggerList	[]string	`json:"reportingTriggerList,omitempty"`
	 ReportIntervalNr	string	`json:"reportIntervalNr,omitempty"`
	 ReportType	string	`json:"reportType,omitempty"`
	 LoggingIntervalNr	string	`json:"loggingIntervalNr,omitempty"`
	 AddPositioningMethodList	[]string	`json:"addPositioningMethodList,omitempty"`
	 MeasurementLteList	[]string	`json:"measurementLteList,omitempty"`
	 ReportInterval	string	`json:"reportInterval,omitempty"`
	 MeasurementPeriodLte	string	`json:"measurementPeriodLte,omitempty"`
	 MbsfnAreaList	[]MbsfnArea	`json:"mbsfnAreaList,omitempty"`
	 EventThresholdRsrp	*int	`json:"eventThresholdRsrp,omitempty"`
	 MdtAllowedPlmnIdList	[]PlmnId	`json:"mdtAllowedPlmnIdList,omitempty"`
}
