package domain

type Response struct {
	//variable type   notation
	Mensaje    string `json:"mensaje"`
}

type Location struct {
	Status string `json:status`
	Data struct {
		Ipv4				string `json:"ipv4,omitempty"`
		Continent_name		string `json:"continent_name,omitempty"`
		Country_name		string `json:"country_name,omitempty"`
		Subdivision_1_name	string `json:"subdivision_1_name,omitempty"`
		Subdivision_2_name	string `json:"subdivision_2_name,omitempty"`
		City_name			string `json:"city_name,omitempty"`
		Latitude			string `json:"latitude"`
		Longitude			string `json:"longitude"`
	} `json:data`
}