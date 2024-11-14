package models
type UpdateSmContextResponse struct {
	 JsonData	*SmContextUpdatedData	`json:"jsonData,omitempty"`
	 BinaryDataN1SmMessage	[]byte	`json:"binaryDataN1SmMessage,omitempty"`
	 BinaryDataN2SmInformation	[]byte	`json:"binaryDataN2SmInformation,omitempty"`
}
