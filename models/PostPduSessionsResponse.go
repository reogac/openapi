package models

type PostPduSessionsResponse struct {
	JsonData               *PduSessionCreatedData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoToUe []byte                 `json:"binaryDataN1SmInfoToUe,omitempty"`
}
