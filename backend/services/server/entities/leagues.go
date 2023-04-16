package entities

type Leagues struct {
	Leagues []League `json:"leagues"`
}

type League struct {
	LeagueName string `json:"league_name"`
	Active bool `json:"active"`
}