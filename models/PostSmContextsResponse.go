package models

type PostSmContextsResponse struct {
	JsonData                  *SmContextCreatedData `json:"jsonData,omitempty"`
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty"`
}
