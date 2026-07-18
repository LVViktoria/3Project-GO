package notifier

import (
	"bytes"
	"fmt"
	"html/template"
	"jobNotifier/internal/model"
	"net/smtp"
	"strings"
	"time"
)

func SendEmail(jobs []model.Job, smtpConfig model.SMTPConfig, recipient string) error {
	if len(jobs) == 0 {
		fmt.Println("No jobs to send")
		return nil
	}

	//создается тело письма
	body, err := buildEmailBody(jobs)
	if err != nil {
		return fmt.Errorf("Error building email body: %s", err)
	}

	//настройка smtp
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	to := []string{recipient}
	msg := buildEmailMessage(smtpConfig.From, to, body)

	addr := fmt.Sprintf("%s:%d", smtpConfig.Host, smtpConfig.Port)
	err = smtp.SendMail(addr, auth, smtpConfig.From, to, msg)
	if err != nil {
		return fmt.Errorf("ошибка отправки письма: %w", err)
	}

	fmt.Printf("Письмо отправлено на %s (%d вакансий)\n", recipient, len(jobs))
	return nil
}

func buildEmailBody(jobs []model.Job) (string, error) {
	// Шаблон письма (вынести в отдельный файл)
	const emailTemplate = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #4CAF50; color: white; padding: 15px; text-align: center; border-radius: 5px; }
        .job { border: 1px solid #ddd; margin: 15px 0; padding: 15px; border-radius: 5px; }
        .job-title { font-size: 18px; font-weight: bold; color: #333; }
        .job-company { color: #666; margin: 5px 0; }
        .job-salary { color: #4CAF50; font-weight: bold; }
        .job-location { color: #666; }
        .job-source { color: #999; font-size: 12px; }
        .job-link { display: inline-block; margin-top: 10px; color: #4CAF50; text-decoration: none; }
        .footer { margin-top: 30px; text-align: center; color: #999; font-size: 12px; }
    </style>
</head>
<body>
    <div class="header">
        <h2> Новые вакансии ({{.Count}})</h2>
    </div>
    
    {{range .Jobs}}
    <div class="job">
        <div class="job-title">{{.Title}}</div>
        <div class="job-company"> {{.Company}}</div>
        <div class="job-salary">{{.Salary}}</div>
        <div class="job-location">{{.Location}}</div>
        <div class="job-source">{{.Source}}</div>
        <a href="{{.URL}}" class="job-link">Подробнее</a>
    </div>
    {{end}}
    
    <div class="footer">
        <p>Job Notifier | {{.Date}}</p>
        <p>Вы получили это письмо, потому что подписаны на уведомления о вакансиях.</p>
    </div>
</body>
</html>
`
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return "", fmt.Errorf("Error parsing email template: %w", err)
	}

	//данные для шаблона-вынести отдельно?
	data := struct {
		Count int
		Jobs  []model.Job
		Date  string
	}{
		Count: len(jobs),
		Jobs:  jobs,
		Date:  time.Now().Format("2006-01-02 15:04:05"), //исправить формат
	}

	//заполнение шаблона
	var buf bytes.Buffer //что за буфер?
	if err = tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("Error executing email template: %w", err)
	}
	return buf.String(), nil
}

// buildEmailMessage собирает полное письмо с заголовками
func buildEmailMessage(from string, to []string, body string) []byte {
	var jobs []model.Job
	headers := map[string]string{
		"From":         from,
		"To":           strings.Join(to, ","),
		"Subject":      fmt.Sprintf("%d новых вакансий по вашему запросу", len(jobs)),
		"MIME-Version": "1.0",
		"Content-Type": `text/html; charset="UTF-8"`,
	}

	//сборка письма
	var buf bytes.Buffer
	for k, v := range headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	buf.WriteString("\r\n")
	buf.WriteString(body)

	return buf.Bytes()
}
