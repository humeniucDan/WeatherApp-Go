package models

type WeatherData struct {
	CurrentWeather struct {
		Temperature   float64 `json:"temperature"`
		WindSpeed     float64 `json:"windspeed"`
		WindDirection float64 `json:"winddirection"`
		Humidity      int     `json:"humidity"`
		Pressure      float64 `json:"pressure_msl"`
		Visibility    float64 `json:"visibility"`
		UVIndex       float64 `json:"uv_index"`
	} `json:"current_weather"`
	Daily struct {
		TemperatureMin []float64 `json:"temperature_2m_min"`
		TemperatureMax []float64 `json:"temperature_2m_max"`
		Sunrise        []string  `json:"sunrise"`
		Sunset         []string  `json:"sunset"`
	} `json:"daily"`
	Hourly struct {
		Temperature2m []float64 `json:"temperature_2m"`
		Time          []string  `json:"time"`
		Precipitation []float64 `json:"precipitation"`
	} `json:"hourly"`
}
