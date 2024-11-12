package models

type EasIpReplacementInfo struct {
	Source EasServerAddress `json:"source"`
	Target EasServerAddress `json:"target"`
}
