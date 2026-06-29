package main

import (
	"flag"
	"fmt"
	"log"
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
		showHistory()
		return
	}

	if *clearHistory {
		if err := history.Clear(); err != nil {
			log.Fatal("Ошибка очистки истории:", err)
		}
		fmt.Println("История очищена")
		return
	}

	var err error
	var geoData *ipdetector.GeoData

	//определение города

	if *city != "" {
		geoData, err = ipdetector.GetLocation(*city, cfg)
	} else {
		cityIP, err := ipdetector.GetCityIP()
		if err != nil {
			log.Fatal("Не удалось определить город по IP", err)
		}
		geoData = &ipdetector.GeoData{City: cityIP}
	}

	if err != nil {
		log.Fatal("Ошибка определения геопозиции:", err)
	}
	if geoData == nil {
		log.Fatal("Не удалось получить геоданные:")
	}

	//получение погоды
	weatherData, err := weather.GetWeather(*geoData, *format, cfg)
	if err != nil {
		log.Fatal("Ошибка получения погоды:", err)
	}
	//вывод погоды
	fmt.Println(weather.Format(*weatherData))

	//сохранение в историю
	entry := history.Entry{
		City:        *city,
		Temperature: weatherData.Temperature,
		Time:        time.Now(),
	}
	if err := history.Save(entry); err != nil {
		fmt.Println("Не удалось сохранить историю", err)
	}
}

func showHistory() {
	entries, err := history.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки истории", err)
	}
	if len(entries) == 0 {
		fmt.Println("История пуста")
	}

	for _, entry := range entries {
		fmt.Printf(
			"%s | %s | %s\n",
			entry.Time.Format("02.01.2006 15:04"),
			entry.City,
			entry.Temperature,
		)
	}
}
