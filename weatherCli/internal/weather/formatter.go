package weather

import "fmt"

func Format(data WeatherData) string {
	return fmt.Sprintf(
		"Город: %s\nПогода: %s",
		data.City,
		data.Temperature,
	)
}
