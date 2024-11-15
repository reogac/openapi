/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	Dnn                        string             `json:"dnn"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
}
