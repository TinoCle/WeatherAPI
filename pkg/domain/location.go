package domain

type LocationSearch struct {
	//variable       type               notation
	Name string `json:"display_name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}
