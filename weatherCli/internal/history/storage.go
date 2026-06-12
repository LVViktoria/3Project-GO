package history

import (
	"encoding/json"
	"os"
)

const HistoryFile = "data/history.json"

func Save(entry Entry) error {
	var entries []Entry

	data, err := os.ReadFile(HistoryFile)
	if err == nil {
		_ = json.Unmarshal(data, &entries)
	}

	entries = append(entries, entry)

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

	var entries []Entry

	if len(data) == 0 {
		return []Entry{}, nil
	}

	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	return entries, nil

}
