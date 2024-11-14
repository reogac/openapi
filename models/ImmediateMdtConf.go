package models
type ImmediateMdtConf struct {
	 AreaScope	*AreaScope	`json:"areaScope,omitempty"`
	 EventThresholdRsrp	*int	`json:"eventThresholdRsrp,omitempty"`
	 EventThresholdRsrq	*int	`json:"eventThresholdRsrq,omitempty"`
	 EventThresholdRsrpNr	*int	`json:"eventThresholdRsrpNr,omitempty"`
	 EventThresholdRsrqNr	*int	`json:"eventThresholdRsrqNr,omitempty"`
	 AddPositioningMethodList	[]string	`json:"addPositioningMethodList,omitempty"`
	 MeasurementLteList	[]string	`json:"measurementLteList,omitempty"`
	 ReportingTriggerList	[]string	`json:"reportingTriggerList,omitempty"`
	 ReportInterval	ReportIntervalMdt	`json:"reportInterval,omitempty"`
	 MdtAllowedPlmnIdList	[]PlmnId	`json:"mdtAllowedPlmnIdList,omitempty"`
	 CollectionPeriodRmmNr	CollectionPeriodRmmNrMdt	`json:"collectionPeriodRmmNr,omitempty"`
	 MeasurementPeriodLte	MeasurementPeriodLteMdt	`json:"measurementPeriodLte,omitempty"`
	 PositioningMethod	PositioningMethodMdt	`json:"positioningMethod,omitempty"`
	 SensorMeasurementList	[]string	`json:"sensorMeasurementList,omitempty"`
	 ReportIntervalNr	ReportIntervalNrMdt	`json:"reportIntervalNr,omitempty"`
	 ReportAmount	ReportAmountMdt	`json:"reportAmount,omitempty"`
	 CollectionPeriodRmmLte	CollectionPeriodRmmLteMdt	`json:"collectionPeriodRmmLte,omitempty"`
	 JobType	JobType	`json:"jobType"`
	 MeasurementNrList	[]string	`json:"measurementNrList,omitempty"`
}
