package models
type MdtConfiguration struct {
	 ReportIntervalNr	ReportIntervalNrMdt	`json:"reportIntervalNr,omitempty"`
	 EventThresholdRsrq	*int	`json:"eventThresholdRsrq,omitempty"`
	 LoggingInterval	LoggingIntervalMdt	`json:"loggingInterval,omitempty"`
	 AreaScope	*AreaScope	`json:"areaScope,omitempty"`
	 ReportInterval	ReportIntervalMdt	`json:"reportInterval,omitempty"`
	 EventThresholdRsrpNr	*int	`json:"eventThresholdRsrpNr,omitempty"`
	 MeasurementPeriodLte	MeasurementPeriodLteMdt	`json:"measurementPeriodLte,omitempty"`
	 ReportingTriggerList	[]string	`json:"reportingTriggerList,omitempty"`
	 SensorMeasurementList	[]string	`json:"sensorMeasurementList,omitempty"`
	 LoggingDuration	LoggingDurationMdt	`json:"loggingDuration,omitempty"`
	 CollectionPeriodRmmLte	CollectionPeriodRmmLteMdt	`json:"collectionPeriodRmmLte,omitempty"`
	 MbsfnAreaList	[]MbsfnArea	`json:"mbsfnAreaList,omitempty"`
	 ReportType	ReportTypeMdt	`json:"reportType,omitempty"`
	 LoggingIntervalNr	LoggingIntervalNrMdt	`json:"loggingIntervalNr,omitempty"`
	 JobType	JobType	`json:"jobType"`
	 MeasurementNrList	[]string	`json:"measurementNrList,omitempty"`
	 EventThresholdRsrqNr	*int	`json:"eventThresholdRsrqNr,omitempty"`
	 AddPositioningMethodList	[]string	`json:"addPositioningMethodList,omitempty"`
	 PositioningMethod	PositioningMethodMdt	`json:"positioningMethod,omitempty"`
	 InterFreqTargetList	[]InterFreqTargetInfo	`json:"interFreqTargetList,omitempty"`
	 ReportAmount	ReportAmountMdt	`json:"reportAmount,omitempty"`
	 MeasurementLteList	[]string	`json:"measurementLteList,omitempty"`
	 EventThresholdRsrp	*int	`json:"eventThresholdRsrp,omitempty"`
	 EventList	[]string	`json:"eventList,omitempty"`
	 LoggingDurationNr	LoggingDurationNrMdt	`json:"loggingDurationNr,omitempty"`
	 CollectionPeriodRmmNr	CollectionPeriodRmmNrMdt	`json:"collectionPeriodRmmNr,omitempty"`
	 MdtAllowedPlmnIdList	[]PlmnId	`json:"mdtAllowedPlmnIdList,omitempty"`
}
