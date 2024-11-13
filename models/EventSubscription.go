package models

type EventSubscription struct {
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
}
