package models

type UpdateSmContextRequest struct {
	BinaryDataN2SmInformationExt1 []byte               `json:"binaryDataN2SmInformationExt1,omitempty"`
	JsonData                      *SmContextUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage         []byte               `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation     []byte               `json:"binaryDataN2SmInformation,omitempty"`
}
