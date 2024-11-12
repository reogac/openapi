package models

type Polygon struct {
	PointList []GeographicalCoordinates `json:"pointList"`
	Shape     SupportedGADShapes        `json:"shape"`
}
