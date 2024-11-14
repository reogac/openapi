package models
type EpsPdnCnxInfo struct {
	 PgwS8cFteid	string	`json:"pgwS8cFteid"`
	 PgwNodeName	string	`json:"pgwNodeName,omitempty"`
	 LinkedBearerId	*int	`json:"linkedBearerId,omitempty"`
}
