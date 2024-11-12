package models

type EventSubscription struct {
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
}
