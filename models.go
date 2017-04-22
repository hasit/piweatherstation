package main

type Weather struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Offset    int     `json:"offset"`
	Currently struct {
		Time                 int     `json:"time"`
		Summary              string  `json:"summary"`
		Icon                 string  `json:"icon"`
		NearestStormDistance int     `json:"nearestStormDistance"`
		NearestStormBearing  int     `json:"nearestStormBearing"`
		PrecipIntensity      int     `json:"precipIntensity"`
		PrecipProbability    int     `json:"precipProbability"`
		Temperature          float64 `json:"temperature"`
		ApparentTemperature  float64 `json:"apparentTemperature"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		WindSpeed            float64 `json:"windSpeed"`
		WindBearing          int     `json:"windBearing"`
		Visibility           int     `json:"visibility"`
		CloudCover           float64 `json:"cloudCover"`
		Pressure             float64 `json:"pressure"`
		Ozone                float64 `json:"ozone"`
	} `json:"currently"`
}
