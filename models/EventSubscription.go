package models

type EventSubscription struct {
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	Dnais              []string                      `json:"dnais,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
}
