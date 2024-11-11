package models

type MoExpDataCounter struct {
	Counter   int    `json:"counter"`
	TimeStamp string `json:"timeStamp,omitempty"`
}
