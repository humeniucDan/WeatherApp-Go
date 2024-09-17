package main

import (
	"fmt"
	"go-lang/services"
)

func main() {
	data, _ := services.GetWeather(46.7667, 23.36)

	ora := 22
	fmt.Println(data.Hourly.Temperature2m[ora], " grade la ora", data.Hourly.Time[ora],
		" si ", data.Hourly.Precipitation[ora]*100, "\b% sanse de ploaie")
	fmt.Println(data.CurrentWeather)
}
