package main

import (
	"flag"
	"fmt"
	"time"
	"weatherCli/internal/history"
	"weatherCli/internal/ipdetector"
	"weatherCli/internal/weather"
	"weatherCli/pkg/config"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	historyFlag := flag.Bool("history", false, "Показать историю запросов")
	clearHistory := flag.Bool("clear-history", false, "Очистить историю")
	cfg := config.Load()

	flag.Parse()

	if *historyFlag {
		entries, err := history.Load()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, entry := range entries {
			fmt.Printf(
				"%s | %s | %s\n",
				entry.Time.Format("02.01.2006 15:04"),
				entry.City,
				entry.Temperature,
			)
		}

		return
	}
	if *clearHistory {
		err := history.Clear()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("История очищена")
		return
	}

	fmt.Println(*city)

	geoData, err := ipdetector.GetLocation(*city, cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)

	weatherData, err := weather.GetWeather(*geoData, *format, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(weatherData)
	err = history.Save(history.Entry{
		City:        geoData.City,
		Temperature: weatherData.Temperature,
		Time:        time.Now(),
	})

	if err != nil {
		fmt.Println(err)
	}
	entries, err := history.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, e := range entries {
		fmt.Println(e.Time, e.City, e.Temperature)
	}
}
