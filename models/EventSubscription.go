package models

type EventSubscription struct {
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	ExptAnaType        string                        `json:"exptAnaType,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NotificationMethod string                        `json:"notificationMethod,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	Event              string                        `json:"event"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	MatchingDir        string                        `json:"matchingDir,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
}
