package models
type PostPduSessionsErrorResponse struct {
	 JsonData	*PduSessionCreateError	`json:"jsonData,omitempty"`
	 BinaryDataN1SmInfoToUe	[]byte	`json:"binaryDataN1SmInfoToUe,omitempty"`
}
