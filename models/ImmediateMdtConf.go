package models

type ImmediateMdtConf struct {
	EventThresholdRsrqNr     *int       `json:"eventThresholdRsrqNr,omitempty"`
	CollectionPeriodRmmLte   string     `json:"collectionPeriodRmmLte,omitempty"`
	MeasurementPeriodLte     string     `json:"measurementPeriodLte,omitempty"`
	AreaScope                *AreaScope `json:"areaScope,omitempty"`
	JobType                  string     `json:"jobType"`
	MeasurementLteList       []string   `json:"measurementLteList,omitempty"`
	ReportIntervalNr         string     `json:"reportIntervalNr,omitempty"`
	EventThresholdRsrp       *int       `json:"eventThresholdRsrp,omitempty"`
	PositioningMethod        string     `json:"positioningMethod,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId   `json:"mdtAllowedPlmnIdList,omitempty"`
	SensorMeasurementList    []string   `json:"sensorMeasurementList,omitempty"`
	MeasurementNrList        []string   `json:"measurementNrList,omitempty"`
	ReportingTriggerList     []string   `json:"reportingTriggerList,omitempty"`
	CollectionPeriodRmmNr    string     `json:"collectionPeriodRmmNr,omitempty"`
	AddPositioningMethodList []string   `json:"addPositioningMethodList,omitempty"`
	ReportInterval           string     `json:"reportInterval,omitempty"`
	ReportAmount             string     `json:"reportAmount,omitempty"`
	EventThresholdRsrpNr     *int       `json:"eventThresholdRsrpNr,omitempty"`
	EventThresholdRsrq       *int       `json:"eventThresholdRsrq,omitempty"`
}
