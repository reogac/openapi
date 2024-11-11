package models
type V2xContext struct {
	 Pc5QoSPara	*Pc5QoSPara	`json:"pc5QoSPara,omitempty"`
	 NrV2xServicesAuth	*NrV2xAuth	`json:"nrV2xServicesAuth,omitempty"`
	 LteV2xServicesAuth	*LteV2xAuth	`json:"lteV2xServicesAuth,omitempty"`
	 NrUeSidelinkAmbr	string	`json:"nrUeSidelinkAmbr,omitempty"`
	 LteUeSidelinkAmbr	string	`json:"lteUeSidelinkAmbr,omitempty"`
}
