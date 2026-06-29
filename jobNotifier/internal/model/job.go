package model

type Job struct {
	//ID      string `json:"id"`
	Title    string `json:"title"`
	Company  string `json:"company"`
	URL      string `json:"url"`
	Salary   string `json:"salary"`
	Location string `json:"location"`
	Source   string `json:"source"`
}

/*type UserSettings struct {
	Keywords  []string `yaml:"keywords"`
	City      string   `yaml:"city"`
	Country   string   `yaml:"country"`
	MinSalary int      `yaml:"min_salary"`
}*/
