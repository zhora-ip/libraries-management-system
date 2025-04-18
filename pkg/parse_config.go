package pkg

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ParseConfig(dest interface{}, path string) {
	if fileExists(path) {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		if err = yaml.Unmarshal(data, dest); err != nil {
			log.Fatal(err)
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
