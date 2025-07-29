package translator

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

type Language string

const (
	EN Language = "en"
	RU Language = "ru"
)

var (
	translations = make(map[string]map[Language]string)
	mu           sync.RWMutex
)

// LoadTranslations загружает переводы из JSON-файла
func LoadTranslations(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var data map[string]map[string]string
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()
	translations = make(map[string]map[Language]string)
	for key, langs := range data {
		translations[key] = make(map[Language]string)
		for lang, val := range langs {
			translations[key][Language(lang)] = val
		}
	}
	return nil
}

// Translate возвращает перевод по ключу и языку
func Translate(key string, lang Language) string {
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := translations[key]; ok {
		if tr, ok := val[lang]; ok {
			return tr
		}
		if tr, ok := val[EN]; ok {
			return tr
		}
	}
	return key
}
