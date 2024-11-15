/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	Dnn                        string             `json:"dnn"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
}
