package models

type PostSmContextsErrorResponse struct {
	JsonData              *SmContextCreateError `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage []byte                `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmMessage []byte                `json:"binaryDataN2SmMessage,omitempty"`
}
