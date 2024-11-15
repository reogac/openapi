/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	Dnn                        string             `json:"dnn"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
}
