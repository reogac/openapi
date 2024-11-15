/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:25 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
}
