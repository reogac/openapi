package models
type NonUeN2MessageTransferRequest struct {
	 BinaryDataN2Information	[]byte	`json:"binaryDataN2Information,omitempty"`
	 JsonData	*N2InformationTransferReqData	`json:"jsonData,omitempty"`
}
