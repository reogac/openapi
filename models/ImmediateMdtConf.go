package models

type ImmediateMdtConf struct {
	EventThresholdRsrp       *int       `json:"eventThresholdRsrp,omitempty"`
	EventThresholdRsrq       *int       `json:"eventThresholdRsrq,omitempty"`
	AreaScope                *AreaScope `json:"areaScope,omitempty"`
	MeasurementPeriodLte     string     `json:"measurementPeriodLte,omitempty"`
	ReportingTriggerList     []string   `json:"reportingTriggerList,omitempty"`
	ReportIntervalNr         string     `json:"reportIntervalNr,omitempty"`
	ReportAmount             string     `json:"reportAmount,omitempty"`
	CollectionPeriodRmmLte   string     `json:"collectionPeriodRmmLte,omitempty"`
	EventThresholdRsrqNr     *int       `json:"eventThresholdRsrqNr,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId   `json:"mdtAllowedPlmnIdList,omitempty"`
	SensorMeasurementList    []string   `json:"sensorMeasurementList,omitempty"`
	EventThresholdRsrpNr     *int       `json:"eventThresholdRsrpNr,omitempty"`
	CollectionPeriodRmmNr    string     `json:"collectionPeriodRmmNr,omitempty"`
	PositioningMethod        string     `json:"positioningMethod,omitempty"`
	AddPositioningMethodList []string   `json:"addPositioningMethodList,omitempty"`
	JobType                  string     `json:"jobType"`
	MeasurementLteList       []string   `json:"measurementLteList,omitempty"`
	MeasurementNrList        []string   `json:"measurementNrList,omitempty"`
	ReportInterval           string     `json:"reportInterval,omitempty"`
}
