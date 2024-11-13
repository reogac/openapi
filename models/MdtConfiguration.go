package models

type MdtConfiguration struct {
	MeasurementNrList        []string                  `json:"measurementNrList,omitempty"`
	ReportInterval           ReportIntervalMdt         `json:"reportInterval,omitempty"`
	EventThresholdRsrq       *int                      `json:"eventThresholdRsrq,omitempty"`
	MeasurementPeriodLte     MeasurementPeriodLteMdt   `json:"measurementPeriodLte,omitempty"`
	LoggingDuration          LoggingDurationMdt        `json:"loggingDuration,omitempty"`
	LoggingDurationNr        LoggingDurationNrMdt      `json:"loggingDurationNr,omitempty"`
	InterFreqTargetList      []InterFreqTargetInfo     `json:"interFreqTargetList,omitempty"`
	EventThresholdRsrpNr     *int                      `json:"eventThresholdRsrpNr,omitempty"`
	LoggingIntervalNr        LoggingIntervalNrMdt      `json:"loggingIntervalNr,omitempty"`
	MbsfnAreaList            []MbsfnArea               `json:"mbsfnAreaList,omitempty"`
	AreaScope                *AreaScope                `json:"areaScope,omitempty"`
	ReportingTriggerList     []string                  `json:"reportingTriggerList,omitempty"`
	CollectionPeriodRmmLte   CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	MeasurementLteList       []string                  `json:"measurementLteList,omitempty"`
	EventThresholdRsrp       *int                      `json:"eventThresholdRsrp,omitempty"`
	EventList                []string                  `json:"eventList,omitempty"`
	PositioningMethod        PositioningMethodMdt      `json:"positioningMethod,omitempty"`
	AddPositioningMethodList []string                  `json:"addPositioningMethodList,omitempty"`
	JobType                  JobType                   `json:"jobType"`
	ReportType               ReportTypeMdt             `json:"reportType,omitempty"`
	EventThresholdRsrqNr     *int                      `json:"eventThresholdRsrqNr,omitempty"`
	SensorMeasurementList    []string                  `json:"sensorMeasurementList,omitempty"`
	ReportIntervalNr         ReportIntervalNrMdt       `json:"reportIntervalNr,omitempty"`
	LoggingInterval          LoggingIntervalMdt        `json:"loggingInterval,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId                  `json:"mdtAllowedPlmnIdList,omitempty"`
	ReportAmount             ReportAmountMdt           `json:"reportAmount,omitempty"`
	CollectionPeriodRmmNr    CollectionPeriodRmmNrMdt  `json:"collectionPeriodRmmNr,omitempty"`
}
