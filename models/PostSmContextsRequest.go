package models

type PostSmContextsRequest struct {
	JsonData                      *SmContextCreateData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage         []byte               `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation     []byte               `json:"binaryDataN2SmInformation,omitempty"`
	BinaryDataN2SmInformationExt1 []byte               `json:"binaryDataN2SmInformationExt1,omitempty"`
}