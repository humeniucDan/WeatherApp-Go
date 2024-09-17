package services

import (
	"encoding/json"
	"fmt"
	"github.com/innotechdevops/openmeteo"
	"go-lang/models"
)

func GetWeather(lat float32, lng float32) (models.WeatherData, error) {
	param := openmeteo.Parameter{
		Latitude:  openmeteo.Float32(lat),
		Longitude: openmeteo.Float32(lng),
		Hourly: &[]string{
			openmeteo.HourlyTemperature2m,
			openmeteo.HourlyWindSpeed10m,
			openmeteo.HourlyVisibility,
			openmeteo.HourlyPrecipitation,
		},
		Daily: &[]string{
			openmeteo.DailyTemperature2mMin,
			openmeteo.DailyTemperature2mMax,
			openmeteo.DailySunrise,
			openmeteo.DailySunset,
			openmeteo.DailyUvIndexMax,
		},
		CurrentWeather: openmeteo.Bool(true),
		Timezone:       openmeteo.String("auto"),
	}

	m := openmeteo.New()
	resp, err := m.Execute(param)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	var data models.WeatherData
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		panic(err)
	}

	return data, nil
}
