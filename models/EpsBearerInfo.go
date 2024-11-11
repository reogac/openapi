package models

type EpsBearerInfo struct {
	PgwS8uFteid    string `json:"pgwS8uFteid"`
	BearerLevelQoS string `json:"bearerLevelQoS"`
	Ebi            int    `json:"ebi"`
}
