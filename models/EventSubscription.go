package models

type EventSubscription struct {
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
}
