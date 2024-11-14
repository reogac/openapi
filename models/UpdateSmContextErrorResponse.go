package models
type UpdateSmContextErrorResponse struct {
	 JsonData	*SmContextUpdateError	`json:"jsonData,omitempty"`
	 BinaryDataN1SmMessage	[]byte	`json:"binaryDataN1SmMessage,omitempty"`
	 BinaryDataN2SmInformation	[]byte	`json:"binaryDataN2SmInformation,omitempty"`
}
