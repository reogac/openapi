package models
type CreateUEContextErrorResponse struct {
	 BinaryDataN2Information	[]byte	`json:"binaryDataN2Information,omitempty"`
	 JsonData	*UeContextCreateError	`json:"jsonData,omitempty"`
}
