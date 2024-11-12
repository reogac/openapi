package models

type ReleaseSmContextRequest struct {
	JsonData                  *SmContextReleaseData `json:"jsonData,omitempty"`
	BinaryDataN2SmInformation string                `json:"binaryDataN2SmInformation,omitempty"`
}
