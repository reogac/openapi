package models

type ImmediateMdtConf struct {
	MeasurementNrList        []string                  `json:"measurementNrList,omitempty"`
	CollectionPeriodRmmLte   CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	MeasurementLteList       []string                  `json:"measurementLteList,omitempty"`
	MeasurementPeriodLte     MeasurementPeriodLteMdt   `json:"measurementPeriodLte,omitempty"`
	AddPositioningMethodList []string                  `json:"addPositioningMethodList,omitempty"`
	EventThresholdRsrpNr     *int                      `json:"eventThresholdRsrpNr,omitempty"`
	ReportIntervalNr         ReportIntervalNrMdt       `json:"reportIntervalNr,omitempty"`
	ReportAmount             ReportAmountMdt           `json:"reportAmount,omitempty"`
	EventThresholdRsrp       *int                      `json:"eventThresholdRsrp,omitempty"`
	EventThresholdRsrq       *int                      `json:"eventThresholdRsrq,omitempty"`
	AreaScope                *AreaScope                `json:"areaScope,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId                  `json:"mdtAllowedPlmnIdList,omitempty"`
	ReportingTriggerList     []string                  `json:"reportingTriggerList,omitempty"`
	ReportInterval           ReportIntervalMdt         `json:"reportInterval,omitempty"`
	EventThresholdRsrqNr     *int                      `json:"eventThresholdRsrqNr,omitempty"`
	CollectionPeriodRmmNr    CollectionPeriodRmmNrMdt  `json:"collectionPeriodRmmNr,omitempty"`
	PositioningMethod        PositioningMethodMdt      `json:"positioningMethod,omitempty"`
	SensorMeasurementList    []string                  `json:"sensorMeasurementList,omitempty"`
	JobType                  JobType                   `json:"jobType"`
}
