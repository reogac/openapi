package models
type RouteToLocation struct {
	 Dnai	string	`json:"dnai"`
	 RouteInfo	*RouteInformation	`json:"routeInfo,omitempty"`
	 RouteProfId	string	`json:"routeProfId,omitempty"`
}
