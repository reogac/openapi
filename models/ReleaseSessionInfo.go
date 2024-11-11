package models

type ReleaseSessionInfo struct {
	ReleaseCause       string `json:"releaseCause"`
	ReleaseSessionList []int  `json:"releaseSessionList"`
}
