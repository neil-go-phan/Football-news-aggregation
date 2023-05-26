package presenter

type ChartDayResponse struct {
	Hour        int `json:"hour"`
	AmountOfJob int       `json:"amount_of_jobs"`
	Cronjobs []CronjobRunTimes `json:"cronjobs"`
}

type CronjobRunTimes struct {
	Name  string `json:"name"`
	Times int    `json:"times"`
}