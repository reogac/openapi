package models

type ReleaseSmContextRequest struct {
	JsonData                  *SmContextReleaseData `json:"jsonData,omitempty"`
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty"`
}
