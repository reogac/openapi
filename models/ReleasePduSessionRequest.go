package models
type ReleasePduSessionRequest struct {
	 JsonData	*ReleaseData	`json:"jsonData,omitempty"`
	 BinaryDataN4Information	[]byte	`json:"binaryDataN4Information,omitempty"`
	 BinaryDataN4InformationExt1	[]byte	`json:"binaryDataN4InformationExt1,omitempty"`
	 BinaryDataN4InformationExt2	[]byte	`json:"binaryDataN4InformationExt2,omitempty"`
}
