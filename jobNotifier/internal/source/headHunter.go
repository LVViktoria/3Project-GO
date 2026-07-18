package source

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"jobNotifier/internal/model"
	"net/http"
	"net/url"
	"time"
)

type HeadHunterSource struct {
	client *http.Client
}

func NewHeadHunterSource() *HeadHunterSource {
	return &HeadHunterSource{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *HeadHunterSource) Name() string {
	return "hh.ru"
}

func (s *HeadHunterSource) Search(keywords []string) ([]model.Job, error) {

	query := "Golang"
	if len(keywords) > 0 {
		query = keywords[0]
	}

	baseURL := "https://api.hh.ru/vacancies"
	params := url.Values{}
	params.Set("text", query)
	params.Set("per_page", "10")
	params.Set("area", "1")
	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		fullURL,
		nil)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	req.Header.Set("User-Agent", "JobNotifier/1.0 (smtpnotifiertest@gmail.com)")
	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("the server returned an error: %d, body: %s", resp.StatusCode, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	type headHunterResponse struct {
		Items []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Employer struct {
				Name string `json:"name"`
			} `json:"employer"`
			Salary struct {
				From int `json:"from"`
				To   int `json:"to"`
			} `json:"salary"`
			Area struct {
				Name string `json:"name"`
			} `json:"area"`
			AlternateURL string `json:"alternate_url"`
		} `json:"items"`
	}

	var result headHunterResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	jobs := make([]model.Job, 0, len(result.Items))
	for _, item := range result.Items {
		salaryStr := "не указана"
		if item.Salary.From > 0 && item.Salary.To > 0 {
			salaryStr = fmt.Sprintf("%d - %d RUB", item.Salary.From, item.Salary.To)
		} else if item.Salary.From > 0 {
			salaryStr = fmt.Sprintf("от %d RUB", item.Salary.From)
		} else if item.Salary.To > 0 {
			salaryStr = fmt.Sprintf("до %d RUB", item.Salary.To)
		}

		job := model.Job{
			Title:    item.Name,
			Company:  item.Employer.Name,
			URL:      item.AlternateURL,
			Salary:   salaryStr,
			Location: item.Area.Name,
			Source:   s.Name(),
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
