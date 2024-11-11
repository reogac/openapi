package models

type EpsPdnCnxInfo struct {
	PgwNodeName    string `json:"pgwNodeName,omitempty"`
	LinkedBearerId *int   `json:"linkedBearerId,omitempty"`
	PgwS8cFteid    string `json:"pgwS8cFteid"`
}
