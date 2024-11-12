package models

type Polygon struct {
	Shape     string                    `json:"shape"`
	PointList []GeographicalCoordinates `json:"pointList"`
}
