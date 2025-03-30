package storage

import (
	"encoding/json"
	"io"
	"os"

	"github.com/et0/go-todo/internal/models"
)

func Load(t *[]models.Task) {
	file, err := os.OpenFile("internal/storage/tasks.json", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decode := json.NewDecoder(file)
	for {
		if err := decode.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}

func Save(t *[]models.Task) {
	file, err := os.OpenFile("internal/storage/tasks.json", os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encode := json.NewEncoder(file)
	encode.Encode(t)
}
