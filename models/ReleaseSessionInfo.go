package models
type ReleaseSessionInfo struct {
	 ReleaseCause	ReleaseCause	`json:"releaseCause"`
	 ReleaseSessionList	[]int	`json:"releaseSessionList"`
}
