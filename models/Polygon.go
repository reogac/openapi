package models
type Polygon struct {
	 Shape	SupportedGADShapes	`json:"shape"`
	 PointList	[]GeographicalCoordinates	`json:"pointList"`
}
