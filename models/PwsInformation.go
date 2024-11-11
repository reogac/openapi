package models
type PwsInformation struct {
	 PwsContainer	N2InfoContent	`json:"pwsContainer"`
	 BcEmptyAreaList	[]GlobalRanNodeId	`json:"bcEmptyAreaList,omitempty"`
	 SendRanResponse	*bool	`json:"sendRanResponse,omitempty"`
	 OmcId	string	`json:"omcId,omitempty"`
	 MessageIdentifier	int	`json:"messageIdentifier"`
	 SerialNumber	int	`json:"serialNumber"`
}
