package models

type ImmediateMdtConf struct {
	MeasurementNrList        []string   `json:"measurementNrList,omitempty"`
	EventThresholdRsrqNr     *int       `json:"eventThresholdRsrqNr,omitempty"`
	CollectionPeriodRmmLte   string     `json:"collectionPeriodRmmLte,omitempty"`
	ReportAmount             string     `json:"reportAmount,omitempty"`
	EventThresholdRsrpNr     *int       `json:"eventThresholdRsrpNr,omitempty"`
	CollectionPeriodRmmNr    string     `json:"collectionPeriodRmmNr,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId   `json:"mdtAllowedPlmnIdList,omitempty"`
	PositioningMethod        string     `json:"positioningMethod,omitempty"`
	AddPositioningMethodList []string   `json:"addPositioningMethodList,omitempty"`
	SensorMeasurementList    []string   `json:"sensorMeasurementList,omitempty"`
	JobType                  string     `json:"jobType"`
	MeasurementLteList       []string   `json:"measurementLteList,omitempty"`
	ReportIntervalNr         string     `json:"reportIntervalNr,omitempty"`
	EventThresholdRsrq       *int       `json:"eventThresholdRsrq,omitempty"`
	MeasurementPeriodLte     string     `json:"measurementPeriodLte,omitempty"`
	ReportingTriggerList     []string   `json:"reportingTriggerList,omitempty"`
	ReportInterval           string     `json:"reportInterval,omitempty"`
	EventThresholdRsrp       *int       `json:"eventThresholdRsrp,omitempty"`
	AreaScope                *AreaScope `json:"areaScope,omitempty"`
}
