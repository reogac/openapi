package models

type UpdateSmContextRequest struct {
	BinaryDataN2SmInformationExt1 string               `json:"binaryDataN2SmInformationExt1,omitempty"`
	JsonData                      *SmContextUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage         string               `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation     string               `json:"binaryDataN2SmInformation,omitempty"`
}
