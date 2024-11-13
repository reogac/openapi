package models

type EasIpReplacementInfo struct {
	Target EasServerAddress `json:"target"`
	Source EasServerAddress `json:"source"`
}
