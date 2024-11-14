/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AnalyticsSubscription struct {
	NwdafId               string              `json:"nwdafId,omitempty"`
	NwdafSetId            string              `json:"nwdafSetId,omitempty"`
	NwdafSubscriptionList []NwdafSubscription `json:"nwdafSubscriptionList"`
}
