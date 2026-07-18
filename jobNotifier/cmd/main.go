package main

//Todo:
//Создать источники
//Настроить периодический запуск
//Запустить бесконечный цикл
//расширить список ключевых слов
import (
	"context"
	"fmt"
	"jobNotifier/internal/config"
	"jobNotifier/internal/model"
	"jobNotifier/internal/notifier"
	"jobNotifier/internal/source"
	"sync"
	"time"
)

func main() {

	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		fmt.Printf("Ошибка загрузки конфига: %v\n", err)
		return
	}

	fmt.Printf("Уведомления будут отправляться на: %s\n", cfg.User.Email)

	sources := []source.Source{
		source.NewHeadHunterSource(),
	}

	//keywords := []string{"Golang", "Go"} //добавить другие ключевые слова
	//периодический запуск
	interval := time.Duration(cfg.User.CheckInterval) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop() //?
	fmt.Printf("Интервал проверки: %d секунд\n", cfg.User.CheckInterval)

	fmt.Println("\n Первая проверка...")
	runCheck(sources, cfg.User.Keywords, cfg)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Проверка...")
			runCheck(sources, cfg.User.Keywords, cfg)
		}
	}
}
func runCheck(sources []source.Source, keywords []string, cfg *model.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resultsChan := make(chan model.Job, 100) //канал

	var wg sync.WaitGroup //?

	for _, src := range sources {
		wg.Add(1)
		source := src //захват переменной для горутины??

		go func() {

			defer wg.Done() //?

			fmt.Printf("поиск в %s...\n\n", source.Name())
			select {
			case <-ctx.Done():
				fmt.Printf("%s: время вышло\n ", source.Name())
				return
			default:
			}

			jobs, err := source.Search(keywords)
			if err != nil {
				fmt.Printf("%s: ошибка: %v\n", source.Name(), err)
				return
			}

			fmt.Printf("%s: найдено %d вакансий\n", source.Name(), len(jobs))
			for _, job := range jobs {
				select {
				case <-ctx.Done():
					return
				case resultsChan <- job:
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultsChan)
		fmt.Println("Все источники завершили работу")
	}()

	var allJobs []model.Job
	for job := range resultsChan {
		allJobs = append(allJobs, job)
	}
	fmt.Printf("Всего собрано вакансий: %d\n", len(allJobs))

	uniqueJobs := deduplicate(allJobs)
	fmt.Printf("После дедупликации: %d уникальных вакансий\n", len(uniqueJobs))

	if len(uniqueJobs) > 0 {
		if err := notifier.SendEmail(uniqueJobs, cfg.SMTP, cfg.User.Email); err != nil {
			fmt.Printf("Ошибка отправки: %v\n", err)
		} else {
			fmt.Println("Нет новых вакансий для отправки")
		}
	}
}
func deduplicate(jobs []model.Job) []model.Job {
	seen := make(map[string]bool)
	unique := make([]model.Job, 0, len(jobs))
	for _, job := range jobs {
		if !seen[job.URL] {
			seen[job.URL] = true
			unique = append(unique, job)
		}
	}
	return unique
}
