package notifier

import (
	"fmt"
	"jobNotifier/internal/model"
	"strings"
)

func SendEmail(jobs []model.Job, smtpConfig model.SMTPConfig, recipient string) error {
	if len(jobs) == 0 {
		fmt.Println("No jobs to send")
		return nil
	}

	var builder strings.Builder

	builder.WriteString((fmt.Sprint("найдено %d новых вакансий:\n\n", len(jobs))))
	for i, job := range jobs {
		builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, job.Title))
		builder.WriteString(fmt.Sprintf(" Компания: %s\n", job.Company))
		builder.WriteString(fmt.Sprintf(" Зарплата: %s\n", job.Salary))
		builder.WriteString(fmt.Sprintf(" Город: %s\n", job.Location))
		builder.WriteString(fmt.Sprintf(" Источник: %s\n", job.Source))
		builder.WriteString(fmt.Sprintf(" Ссылка: %s\n\n", job.URL))
	}
	message := builder.String()
	fmt.Println("======= ПИСЬМО =======")
	fmt.Println(message)
	fmt.Println("=========================")
	return nil
}
