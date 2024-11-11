package models

type MdtConfiguration struct {
	PositioningMethod        string                `json:"positioningMethod,omitempty"`
	MeasurementPeriodLte     string                `json:"measurementPeriodLte,omitempty"`
	ReportingTriggerList     []string              `json:"reportingTriggerList,omitempty"`
	EventThresholdRsrp       *int                  `json:"eventThresholdRsrp,omitempty"`
	LoggingInterval          string                `json:"loggingInterval,omitempty"`
	CollectionPeriodRmmLte   string                `json:"collectionPeriodRmmLte,omitempty"`
	MbsfnAreaList            []MbsfnArea           `json:"mbsfnAreaList,omitempty"`
	JobType                  string                `json:"jobType"`
	AreaScope                *AreaScope            `json:"areaScope,omitempty"`
	ReportAmount             string                `json:"reportAmount,omitempty"`
	EventThresholdRsrpNr     *int                  `json:"eventThresholdRsrpNr,omitempty"`
	EventThresholdRsrq       *int                  `json:"eventThresholdRsrq,omitempty"`
	EventList                string                `json:"eventList,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId              `json:"mdtAllowedPlmnIdList,omitempty"`
	SensorMeasurementList    []string              `json:"sensorMeasurementList,omitempty"`
	ReportIntervalNr         string                `json:"reportIntervalNr,omitempty"`
	InterFreqTargetList      []InterFreqTargetInfo `json:"interFreqTargetList,omitempty"`
	ReportType               string                `json:"reportType,omitempty"`
	ReportInterval           string                `json:"reportInterval,omitempty"`
	LoggingDuration          string                `json:"loggingDuration,omitempty"`
	EventThresholdRsrqNr     *int                  `json:"eventThresholdRsrqNr,omitempty"`
	CollectionPeriodRmmNr    string                `json:"collectionPeriodRmmNr,omitempty"`
	AddPositioningMethodList []string              `json:"addPositioningMethodList,omitempty"`
	MeasurementLteList       []string              `json:"measurementLteList,omitempty"`
	LoggingDurationNr        string                `json:"loggingDurationNr,omitempty"`
	MeasurementNrList        []string              `json:"measurementNrList,omitempty"`
	LoggingIntervalNr        string                `json:"loggingIntervalNr,omitempty"`
}
