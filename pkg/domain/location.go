package domain

type Search struct {
	//variable       type               notation
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
