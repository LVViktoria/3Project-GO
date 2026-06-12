package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weatherCli/internal/ipdetector"
	"weatherCli/pkg/config"
)

type WeatherData struct {
	City        string
	Temperature string
}

func GetWeather(geo ipdetector.GeoData, format int, cfg config.Config) (*WeatherData, error) {
	baseUrl, err := url.Parse(cfg.WeatherURL + geo.City)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String()) // парсинг, возврат city
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //добавлен defer

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	
	weatherData := &WeatherData{
		City:        geo.City,
		Temperature: string(body),
	}
	return weatherData, nil

}
