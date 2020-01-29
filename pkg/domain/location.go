package domain

type Location struct {
	Status string `json:status`
	Data   struct {
		// Continent_name		string `json:"continent_name,omitempty"`
		// Country_name		string `json:"country_name,omitempty"`
		Subdivision1 string `json:"subdivision_1_name,omitempty"`
		Ciudad       string `json:"city_name,omitempty"`
		Latitud      string `json:"latitude"`
		Longitud     string `json:"longitude"`
	} `json:data`
}

type Search struct {
	Id   string `json:"place_id"`
	Name string `json:"display_name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}

type Locations struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}
