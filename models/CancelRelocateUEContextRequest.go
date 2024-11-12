package models

type CancelRelocateUEContextRequest struct {
	BinaryDataGtpcMessage string                       `json:"binaryDataGtpcMessage,omitempty"`
	JsonData              *UeContextCancelRelocateData `json:"jsonData,omitempty"`
}
