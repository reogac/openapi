package models

type UpdatePduSessionErrorResponse struct {
	JsonData               *HsmfUpdateError `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoToUe []byte           `json:"binaryDataN1SmInfoToUe,omitempty"`
}
