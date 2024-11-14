/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	Supi                    string                    `json:"supi"`
	Dnn                     string                    `json:"dnn"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
}
