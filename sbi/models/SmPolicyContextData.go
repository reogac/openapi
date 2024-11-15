/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Dnn                     string                    `json:"dnn"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	Supi                    string                    `json:"supi"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
}
