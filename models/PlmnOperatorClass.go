package models

type PlmnOperatorClass struct {
	LcsClientIds   []string       `json:"lcsClientIds"`
	LcsClientClass LcsClientClass `json:"lcsClientClass"`
}
