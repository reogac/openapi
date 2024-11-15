/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	Supi                    string                    `json:"supi"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	PduSessionId            int                       `json:"pduSessionId"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	Dnn                     string                    `json:"dnn"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
}
