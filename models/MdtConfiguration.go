package models

type MdtConfiguration struct {
	ReportIntervalNr         ReportIntervalNrMdt       `json:"reportIntervalNr,omitempty"`
	InterFreqTargetList      []InterFreqTargetInfo     `json:"interFreqTargetList,omitempty"`
	SensorMeasurementList    []string                  `json:"sensorMeasurementList,omitempty"`
	EventThresholdRsrqNr     *int                      `json:"eventThresholdRsrqNr,omitempty"`
	CollectionPeriodRmmNr    CollectionPeriodRmmNrMdt  `json:"collectionPeriodRmmNr,omitempty"`
	ReportType               ReportTypeMdt             `json:"reportType,omitempty"`
	EventList                []string                  `json:"eventList,omitempty"`
	LoggingInterval          LoggingIntervalMdt        `json:"loggingInterval,omitempty"`
	AddPositioningMethodList []string                  `json:"addPositioningMethodList,omitempty"`
	MbsfnAreaList            []MbsfnArea               `json:"mbsfnAreaList,omitempty"`
	ReportingTriggerList     []string                  `json:"reportingTriggerList,omitempty"`
	MeasurementLteList       []string                  `json:"measurementLteList,omitempty"`
	MeasurementNrList        []string                  `json:"measurementNrList,omitempty"`
	ReportInterval           ReportIntervalMdt         `json:"reportInterval,omitempty"`
	EventThresholdRsrq       *int                      `json:"eventThresholdRsrq,omitempty"`
	JobType                  JobType                   `json:"jobType"`
	ReportAmount             ReportAmountMdt           `json:"reportAmount,omitempty"`
	PositioningMethod        PositioningMethodMdt      `json:"positioningMethod,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId                  `json:"mdtAllowedPlmnIdList,omitempty"`
	LoggingDurationNr        LoggingDurationNrMdt      `json:"loggingDurationNr,omitempty"`
	CollectionPeriodRmmLte   CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	EventThresholdRsrp       *int                      `json:"eventThresholdRsrp,omitempty"`
	EventThresholdRsrpNr     *int                      `json:"eventThresholdRsrpNr,omitempty"`
	LoggingIntervalNr        LoggingIntervalNrMdt      `json:"loggingIntervalNr,omitempty"`
	LoggingDuration          LoggingDurationMdt        `json:"loggingDuration,omitempty"`
	MeasurementPeriodLte     MeasurementPeriodLteMdt   `json:"measurementPeriodLte,omitempty"`
	AreaScope                *AreaScope                `json:"areaScope,omitempty"`
}
