package models
type ImmediateMdtConf struct {
	 JobType	JobType	`json:"jobType"`
	 MeasurementNrList	[]string	`json:"measurementNrList,omitempty"`
	 ReportInterval	ReportIntervalMdt	`json:"reportInterval,omitempty"`
	 EventThresholdRsrpNr	*int	`json:"eventThresholdRsrpNr,omitempty"`
	 PositioningMethod	PositioningMethodMdt	`json:"positioningMethod,omitempty"`
	 SensorMeasurementList	[]string	`json:"sensorMeasurementList,omitempty"`
	 EventThresholdRsrp	*int	`json:"eventThresholdRsrp,omitempty"`
	 EventThresholdRsrq	*int	`json:"eventThresholdRsrq,omitempty"`
	 CollectionPeriodRmmLte	CollectionPeriodRmmLteMdt	`json:"collectionPeriodRmmLte,omitempty"`
	 MeasurementLteList	[]string	`json:"measurementLteList,omitempty"`
	 EventThresholdRsrqNr	*int	`json:"eventThresholdRsrqNr,omitempty"`
	 MeasurementPeriodLte	MeasurementPeriodLteMdt	`json:"measurementPeriodLte,omitempty"`
	 ReportingTriggerList	[]string	`json:"reportingTriggerList,omitempty"`
	 ReportIntervalNr	ReportIntervalNrMdt	`json:"reportIntervalNr,omitempty"`
	 ReportAmount	ReportAmountMdt	`json:"reportAmount,omitempty"`
	 CollectionPeriodRmmNr	CollectionPeriodRmmNrMdt	`json:"collectionPeriodRmmNr,omitempty"`
	 AreaScope	*AreaScope	`json:"areaScope,omitempty"`
	 AddPositioningMethodList	[]string	`json:"addPositioningMethodList,omitempty"`
	 MdtAllowedPlmnIdList	[]PlmnId	`json:"mdtAllowedPlmnIdList,omitempty"`
}
