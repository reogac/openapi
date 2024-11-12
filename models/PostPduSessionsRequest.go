package models

type PostPduSessionsRequest struct {
	JsonData                  *PduSessionCreateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe  string                `json:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo string                `json:"binaryDataUnknownN1SmInfo,omitempty"`
}
