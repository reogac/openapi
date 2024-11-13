package models

type CreateUEContextErrorResponse struct {
	JsonData                *UeContextCreateError `json:"jsonData,omitempty"`
	BinaryDataN2Information []byte                `json:"binaryDataN2Information,omitempty"`
}
