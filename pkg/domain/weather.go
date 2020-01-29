package domain

type Weather struct {
	Clima []struct {
		Descripcion string `json:"description"`
	} `json:"weather"`
	Detalle struct {
		Temp       		 float64 `json:"temp"`
		SensacionTermica float64 `json:"feels_like"`
		Humedad          int `json:"humidity"`
	} `json:"main"`
	Ciudad string `json:"name"`
}