package models

type EpsPdnCnxInfo struct {
	LinkedBearerId *int   `json:"linkedBearerId,omitempty"`
	PgwS8cFteid    string `json:"pgwS8cFteid"`
	PgwNodeName    string `json:"pgwNodeName,omitempty"`
}
