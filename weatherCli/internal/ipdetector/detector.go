package ipdetector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"weatherCli/pkg/config"
)

type GeoData struct {
	City string `json:"city"`
}

func GetLocation(city string, cfg config.Config) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	resp, err := http.Get(cfg.GeoURL) // парсинг, возврат city
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("NOT200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	if err := json.Unmarshal(body, &geo); err != nil {
		return nil, err
	}
	return &geo, nil
}
func GetCityIP() (string, error) {

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API вернул статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения: %w", err)
	}

	var geo GeoData
	if err := json.Unmarshal(body, &geo); err != nil {
		return "", fmt.Errorf("ошибка парсинга: %w", err)
	}
	return geo.City, nil
}

func GetLocationIP() (string, error) {
	return GetCityIP()
}
