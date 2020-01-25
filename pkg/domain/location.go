package domain

type Search struct {
	//variable       type               notation
	Id   string `json:"place_id"`
	Name string `json:"display_name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}

type Locations struct {
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}

type UpdateLocation struct {
	Id  string `json:"id"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
