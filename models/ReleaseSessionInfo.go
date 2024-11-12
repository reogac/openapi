package models

type ReleaseSessionInfo struct {
	ReleaseSessionList []int  `json:"releaseSessionList"`
	ReleaseCause       string `json:"releaseCause"`
}
