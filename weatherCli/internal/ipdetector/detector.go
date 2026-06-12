package ipdetector

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

	defer resp.Body.Close() //добавлен дефер

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
