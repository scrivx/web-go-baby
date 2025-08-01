package models

type Routes struct {
	MAIN          string
	LOGS          string
	CLOCK         string
	RESTART_CYCLE string
	LOG_TABLE     string
	ERROR         string
	CLEAR_ERROR   string
}

var RoutesInstance = Routes{
	MAIN:          "/",
	LOGS:          "/logs",
	CLOCK:         "/clock",
	RESTART_CYCLE: "/clock/restart-cycle",
	LOG_TABLE:     "/log-table",
	ERROR:         "/error",
	CLEAR_ERROR:   "/clear-error",
}