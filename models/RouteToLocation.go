package models

type RouteToLocation struct {
	RouteInfo   *RouteInformation `json:"routeInfo,omitempty"`
	RouteProfId string            `json:"routeProfId,omitempty"`
	Dnai        string            `json:"dnai"`
}
