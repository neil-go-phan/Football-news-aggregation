package presenter

import "server/services"

type ChartDayResponse struct {
	Hour        int               `json:"hour"`
	AmountOfJob int               `json:"amount_of_jobs"`
	Cronjobs    []CronjobRunTimes `json:"cronjobs"`
}

type CronjobRunTimes struct {
	Name  string `json:"name"`
	Times int    `json:"times"`
}

type ChartHourResponse struct {
	Minute      int                       `json:"minute"`
	AmountOfJob int                       `json:"amount_of_jobs"`
	Cronjobs    []services.CronjobInChart `json:"cronjobs"`
}
