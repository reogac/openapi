package models

type PlmnOperatorClass struct {
	LcsClientClass LcsClientClass `json:"lcsClientClass"`
	LcsClientIds   []string       `json:"lcsClientIds"`
}
