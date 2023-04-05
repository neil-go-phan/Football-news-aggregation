package entities

type Club struct {
	Name string`json:"name"`
	Logo string`json:"logo"`
}

type HtmlClubClass struct {
	Name string `json:"club_name"`
	Logo string `json:"club_logo"`
}