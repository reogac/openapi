/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TscaiInputContainer struct {
	SurTimeInTime    *int   `json:"surTimeInTime,omitempty"`
	Periodicity      *int   `json:"periodicity,omitempty"`
	BurstArrivalTime string `json:"burstArrivalTime,omitempty"`
	SurTimeInNumMsg  *int   `json:"surTimeInNumMsg,omitempty"`
}
