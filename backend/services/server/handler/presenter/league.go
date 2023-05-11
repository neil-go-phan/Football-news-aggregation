package presenter

type League struct {
	LeagueName string `json:"league_name"`
	Active bool `json:"active"`
}

type Leagues struct {
	Leagues []League `json:"leagues"`
}