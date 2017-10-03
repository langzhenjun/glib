package utils

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// LoadJSON
func LoadJSON(c interface{}, path string) error {
	path = strings.Trim(path, " /")

	// if file is empty, do nothing
	if path == "" {
		return os.ErrNotExist
	}

	if !strings.HasSuffix(path, ".json") {
		path = path + ".json"
	}

	file, err := os.Open(path)

	// if open file failed, do nothing with error
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)

	return err
}

// LoadENV
func LoadENV(filename string) error {
	err := godotenv.Load(filename)
	return err
}
