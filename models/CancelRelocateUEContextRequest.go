package models

type CancelRelocateUEContextRequest struct {
	JsonData              *UeContextCancelRelocateData `json:"jsonData,omitempty"`
	BinaryDataGtpcMessage string                       `json:"binaryDataGtpcMessage,omitempty"`
}
