package models

type UpdatePduSessionRequest struct {
	BinaryDataN4InformationExt2 string          `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *HsmfUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe    string          `json:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo   string          `json:"binaryDataUnknownN1SmInfo,omitempty"`
	BinaryDataN4Information     string          `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 string          `json:"binaryDataN4InformationExt1,omitempty"`
}
