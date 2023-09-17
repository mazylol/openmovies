package types

type Movie struct {
	Title       string   `json:"title"`
	Shortname   string   `json:"shortname"`
	Description string   `json:"description"`
	Genre       []string `json:"genre"`
	Rating      string   `json:"rating"`
	Language    string   `json:"language"`
	Cast        []struct {
		Name string `json:"name"`
		Role string `json:"role"`
	} `json:"cast"`
	Crew struct {
		Director []string `json:"director"`
		Producer []string `json:"producer"`
		Writer   []string `json:"writer"`
	} `json:"crew"`
	Distributor []string `json:"distributor"`
	BoxOffice   struct {
		Budget uint `json:"budget"`
		Gross  struct {
			Us        uint `json:"us"`
			Worldwide uint `json:"worldwide"`
		} `json:"gross"`
		OpeningWeekend struct {
			Us        uint `json:"us"`
			Worldwide uint `json:"worldwide"`
		} `json:"openingWeekend"`
	} `json:"boxOffice"`
	Release struct {
		Year  uint `json:"year"`
		Month uint `json:"month"`
		Day   uint `json:"day"`
	} `json:"release"`
	CountryOfOrigin []string `json:"countryOfOrigin"`
	Runtime         struct {
		Hours   uint `json:"hours"`
		Minutes uint `json:"minutes"`
		Seconds uint `json:"seconds"`
	} `json:"runtime"`
}
