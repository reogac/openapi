package models
type ExpectedUeBehavior struct {
	 ExpMoveTrajectory	[]UserLocation	`json:"expMoveTrajectory"`
	 ValidityTime	string	`json:"validityTime"`
}
