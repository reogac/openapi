/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
}
