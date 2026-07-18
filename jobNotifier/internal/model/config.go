package model

type SMTPConfig struct {
	Host     string `yaml:"host"`     // SMTP сервер (например, smtp.gmail.com)
	Port     int    `yaml:"port"`     // Порт (587 для TLS)
	Username string `yaml:"username"` // Логин (email)
	Password string `yaml:"password"` // Пароль или app password
	From     string `yaml:"from"`     // Email отправителя
}

type Config struct {
	SMTP SMTPConfig `yaml:"smtp"`
	User UserConfig `yaml:"user"`
}

type UserConfig struct {
	Keywords      []string `yaml:"keywords"`
	City          string   `yaml:"city"`
	MinSalary     int      `yaml:"min_salary"`
	CheckInterval int      `yaml:"check_interval"`
	Email         string   `yaml:"email"` //Email получателя
}
