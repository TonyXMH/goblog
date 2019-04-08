package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:    "GetAccount",
		Method:  "GET",
		Pattern: "/accounts/{accountID}",
		HandlerFunc:GetAccount,

	},
	Route{
		Name:        "HealthCheck",
		Method:      "GET",
		Pattern:     "/health",
		HandlerFunc: HealthCheck,
	},
	Route{
		Name:        "TestAbility",
		Method:      "GET",
		Pattern:     "/testability/healthy/{state}",
		HandlerFunc: SetHealthyState,
	},
}
