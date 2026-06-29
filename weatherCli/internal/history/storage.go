package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const HistoryFile = "data/history.json"

func Save(entry Entry) error {
	dir := filepath.Dir(HistoryFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("не могу создать папку: %w", err)
	}
	var entries []Entry

	data, err := os.ReadFile(HistoryFile)
	if err == nil {
		if err := json.Unmarshal(data, &entries); err != nil {
			return fmt.Errorf("ошибка парсинга: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	entries = append(entries, entry) //добавление новой записи

	data, err = json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(HistoryFile, data, 0644)
}
func Clear() error {
	return os.WriteFile(HistoryFile, []byte("[]"), 0644)
}
func Load() ([]Entry, error) {
	data, err := os.ReadFile(HistoryFile)
	if os.IsNotExist(err) {
		return []Entry{}, nil
	}
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []Entry{}, nil
	}

	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	return entries, nil

}
