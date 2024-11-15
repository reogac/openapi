/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:28 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
}
