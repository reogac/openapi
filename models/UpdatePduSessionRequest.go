package models

type UpdatePduSessionRequest struct {
	BinaryDataUnknownN1SmInfo   []byte          `json:"binaryDataUnknownN1SmInfo,omitempty"`
	BinaryDataN4Information     []byte          `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 []byte          `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 []byte          `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *HsmfUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe    []byte          `json:"binaryDataN1SmInfoFromUe,omitempty"`
}
