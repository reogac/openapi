package models
type RouteToLocation struct {
	 RouteProfId	string	`json:"routeProfId,omitempty"`
	 Dnai	string	`json:"dnai"`
	 RouteInfo	*RouteInformation	`json:"routeInfo,omitempty"`
}
