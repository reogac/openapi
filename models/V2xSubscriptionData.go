package models
type V2xSubscriptionData struct {
	 LtePc5Ambr	string	`json:"ltePc5Ambr,omitempty"`
	 NrV2xServicesAuth	*NrV2xAuth	`json:"nrV2xServicesAuth,omitempty"`
	 LteV2xServicesAuth	*LteV2xAuth	`json:"lteV2xServicesAuth,omitempty"`
	 NrUePc5Ambr	string	`json:"nrUePc5Ambr,omitempty"`
}
