/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	SmContextRef               string             `json:"smContextRef"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	AccessType                 AccessType         `json:"accessType"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	Dnn                        string             `json:"dnn"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
}
