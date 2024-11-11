package models

type ImmediateMdtConf struct {
	MeasurementNrList        []string   `json:"measurementNrList,omitempty"`
	ReportAmount             string     `json:"reportAmount,omitempty"`
	EventThresholdRsrp       *int       `json:"eventThresholdRsrp,omitempty"`
	JobType                  string     `json:"jobType"`
	ReportingTriggerList     []string   `json:"reportingTriggerList,omitempty"`
	EventThresholdRsrq       *int       `json:"eventThresholdRsrq,omitempty"`
	MeasurementPeriodLte     string     `json:"measurementPeriodLte,omitempty"`
	AreaScope                *AreaScope `json:"areaScope,omitempty"`
	PositioningMethod        string     `json:"positioningMethod,omitempty"`
	AddPositioningMethodList []string   `json:"addPositioningMethodList,omitempty"`
	EventThresholdRsrqNr     *int       `json:"eventThresholdRsrqNr,omitempty"`
	CollectionPeriodRmmNr    string     `json:"collectionPeriodRmmNr,omitempty"`
	SensorMeasurementList    []string   `json:"sensorMeasurementList,omitempty"`
	MeasurementLteList       []string   `json:"measurementLteList,omitempty"`
	ReportInterval           string     `json:"reportInterval,omitempty"`
	ReportIntervalNr         string     `json:"reportIntervalNr,omitempty"`
	EventThresholdRsrpNr     *int       `json:"eventThresholdRsrpNr,omitempty"`
	CollectionPeriodRmmLte   string     `json:"collectionPeriodRmmLte,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId   `json:"mdtAllowedPlmnIdList,omitempty"`
}
