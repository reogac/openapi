package models

type MdtConfiguration struct {
	EventThresholdRsrpNr     *int                      `json:"eventThresholdRsrpNr,omitempty"`
	LoggingDurationNr        LoggingDurationNrMdt      `json:"loggingDurationNr,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId                  `json:"mdtAllowedPlmnIdList,omitempty"`
	LoggingDuration          LoggingDurationMdt        `json:"loggingDuration,omitempty"`
	CollectionPeriodRmmNr    CollectionPeriodRmmNrMdt  `json:"collectionPeriodRmmNr,omitempty"`
	ReportingTriggerList     []string                  `json:"reportingTriggerList,omitempty"`
	EventList                []string                  `json:"eventList,omitempty"`
	JobType                  JobType                   `json:"jobType"`
	ReportType               ReportTypeMdt             `json:"reportType,omitempty"`
	MeasurementLteList       []string                  `json:"measurementLteList,omitempty"`
	InterFreqTargetList      []InterFreqTargetInfo     `json:"interFreqTargetList,omitempty"`
	ReportInterval           ReportIntervalMdt         `json:"reportInterval,omitempty"`
	LoggingInterval          LoggingIntervalMdt        `json:"loggingInterval,omitempty"`
	CollectionPeriodRmmLte   CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	EventThresholdRsrq       *int                      `json:"eventThresholdRsrq,omitempty"`
	LoggingIntervalNr        LoggingIntervalNrMdt      `json:"loggingIntervalNr,omitempty"`
	MeasurementPeriodLte     MeasurementPeriodLteMdt   `json:"measurementPeriodLte,omitempty"`
	MeasurementNrList        []string                  `json:"measurementNrList,omitempty"`
	ReportAmount             ReportAmountMdt           `json:"reportAmount,omitempty"`
	EventThresholdRsrqNr     *int                      `json:"eventThresholdRsrqNr,omitempty"`
	PositioningMethod        PositioningMethodMdt      `json:"positioningMethod,omitempty"`
	AddPositioningMethodList []string                  `json:"addPositioningMethodList,omitempty"`
	EventThresholdRsrp       *int                      `json:"eventThresholdRsrp,omitempty"`
	MbsfnAreaList            []MbsfnArea               `json:"mbsfnAreaList,omitempty"`
	AreaScope                *AreaScope                `json:"areaScope,omitempty"`
	SensorMeasurementList    []string                  `json:"sensorMeasurementList,omitempty"`
	ReportIntervalNr         ReportIntervalNrMdt       `json:"reportIntervalNr,omitempty"`
}
