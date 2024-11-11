package models

type EpsBearerInfo struct {
	BearerLevelQoS string `json:"bearerLevelQoS"`
	Ebi            int    `json:"ebi"`
	PgwS8uFteid    string `json:"pgwS8uFteid"`
}
