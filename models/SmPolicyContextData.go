package models

type SmPolicyContextData struct {
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	Supi                    string                    `json:"supi"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	Dnn                     string                    `json:"dnn"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
}