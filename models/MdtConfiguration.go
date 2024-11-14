package models
type MdtConfiguration struct {
	 MeasurementLteList	[]string	`json:"measurementLteList,omitempty"`
	 ReportInterval	ReportIntervalMdt	`json:"reportInterval,omitempty"`
	 EventThresholdRsrp	*int	`json:"eventThresholdRsrp,omitempty"`
	 EventThresholdRsrq	*int	`json:"eventThresholdRsrq,omitempty"`
	 EventThresholdRsrqNr	*int	`json:"eventThresholdRsrqNr,omitempty"`
	 MeasurementPeriodLte	MeasurementPeriodLteMdt	`json:"measurementPeriodLte,omitempty"`
	 JobType	JobType	`json:"jobType"`
	 EventThresholdRsrpNr	*int	`json:"eventThresholdRsrpNr,omitempty"`
	 LoggingDuration	LoggingDurationMdt	`json:"loggingDuration,omitempty"`
	 CollectionPeriodRmmNr	CollectionPeriodRmmNrMdt	`json:"collectionPeriodRmmNr,omitempty"`
	 InterFreqTargetList	[]InterFreqTargetInfo	`json:"interFreqTargetList,omitempty"`
	 ReportIntervalNr	ReportIntervalNrMdt	`json:"reportIntervalNr,omitempty"`
	 ReportAmount	ReportAmountMdt	`json:"reportAmount,omitempty"`
	 AddPositioningMethodList	[]string	`json:"addPositioningMethodList,omitempty"`
	 CollectionPeriodRmmLte	CollectionPeriodRmmLteMdt	`json:"collectionPeriodRmmLte,omitempty"`
	 MdtAllowedPlmnIdList	[]PlmnId	`json:"mdtAllowedPlmnIdList,omitempty"`
	 MeasurementNrList	[]string	`json:"measurementNrList,omitempty"`
	 EventList	[]string	`json:"eventList,omitempty"`
	 LoggingDurationNr	LoggingDurationNrMdt	`json:"loggingDurationNr,omitempty"`
	 SensorMeasurementList	[]string	`json:"sensorMeasurementList,omitempty"`
	 LoggingIntervalNr	LoggingIntervalNrMdt	`json:"loggingIntervalNr,omitempty"`
	 AreaScope	*AreaScope	`json:"areaScope,omitempty"`
	 ReportingTriggerList	[]string	`json:"reportingTriggerList,omitempty"`
	 LoggingInterval	LoggingIntervalMdt	`json:"loggingInterval,omitempty"`
	 PositioningMethod	PositioningMethodMdt	`json:"positioningMethod,omitempty"`
	 ReportType	ReportTypeMdt	`json:"reportType,omitempty"`
	 MbsfnAreaList	[]MbsfnArea	`json:"mbsfnAreaList,omitempty"`
}
