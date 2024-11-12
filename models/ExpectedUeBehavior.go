package models

type ExpectedUeBehavior struct {
	ValidityTime      string         `json:"validityTime"`
	ExpMoveTrajectory []UserLocation `json:"expMoveTrajectory"`
}
