package models
type PostPduSessionsResponse struct {
	 BinaryDataN1SmInfoToUe	[]byte	`json:"binaryDataN1SmInfoToUe,omitempty"`
	 JsonData	*PduSessionCreatedData	`json:"jsonData,omitempty"`
}
