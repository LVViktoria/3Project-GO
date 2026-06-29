package source

import "jobNotifier/internal/model"

type Source interface {
	Name() string
	Search(keywords []string) ([]model.Job, error)
}
