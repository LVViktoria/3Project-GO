package config

import "time"

type Config struct {
	WeatherURL string
	GeoURL     string
	Timeout    time.Duration //таймаут запроса
}

func Load() Config {
	return Config{
		WeatherURL: "https://wttr.in/",
		GeoURL:     "https://ipapi.co/json/",
		Timeout:    5 * time.Second, //таймаут запроса
	}
}
