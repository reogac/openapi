package models
type UpdatePduSessionResponse struct {
	 JsonData	*HsmfUpdatedData	`json:"jsonData,omitempty"`
	 BinaryDataN1SmInfoToUe	[]byte	`json:"binaryDataN1SmInfoToUe,omitempty"`
	 BinaryDataN4Information	[]byte	`json:"binaryDataN4Information,omitempty"`
	 BinaryDataN4InformationExt1	[]byte	`json:"binaryDataN4InformationExt1,omitempty"`
	 BinaryDataN4InformationExt2	[]byte	`json:"binaryDataN4InformationExt2,omitempty"`
}
