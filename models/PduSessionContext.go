/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SNssai                     Snssai             `json:"sNssai"`
	AccessType                 AccessType         `json:"accessType"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	Dnn                        string             `json:"dnn"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
}
