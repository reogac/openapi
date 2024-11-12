package models

type ReleasePduSessionRequest struct {
	BinaryDataN4InformationExt1 string       `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 string       `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *ReleaseData `json:"jsonData,omitempty"`
	BinaryDataN4Information     string       `json:"binaryDataN4Information,omitempty"`
}
