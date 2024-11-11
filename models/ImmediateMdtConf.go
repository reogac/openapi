package models

type ImmediateMdtConf struct {
	MeasurementLteList       []string   `json:"measurementLteList,omitempty"`
	ReportingTriggerList     []string   `json:"reportingTriggerList,omitempty"`
	ReportIntervalNr         string     `json:"reportIntervalNr,omitempty"`
	MeasurementPeriodLte     string     `json:"measurementPeriodLte,omitempty"`
	SensorMeasurementList    []string   `json:"sensorMeasurementList,omitempty"`
	JobType                  string     `json:"jobType"`
	ReportInterval           string     `json:"reportInterval,omitempty"`
	EventThresholdRsrq       *int       `json:"eventThresholdRsrq,omitempty"`
	AddPositioningMethodList []string   `json:"addPositioningMethodList,omitempty"`
	EventThresholdRsrp       *int       `json:"eventThresholdRsrp,omitempty"`
	EventThresholdRsrpNr     *int       `json:"eventThresholdRsrpNr,omitempty"`
	EventThresholdRsrqNr     *int       `json:"eventThresholdRsrqNr,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId   `json:"mdtAllowedPlmnIdList,omitempty"`
	MeasurementNrList        []string   `json:"measurementNrList,omitempty"`
	ReportAmount             string     `json:"reportAmount,omitempty"`
	CollectionPeriodRmmLte   string     `json:"collectionPeriodRmmLte,omitempty"`
	CollectionPeriodRmmNr    string     `json:"collectionPeriodRmmNr,omitempty"`
	AreaScope                *AreaScope `json:"areaScope,omitempty"`
	PositioningMethod        string     `json:"positioningMethod,omitempty"`
}
