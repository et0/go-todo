package main

import (
	"fmt"
	"os"

	"github.com/et0/go-todo/internal/models"
	"github.com/et0/go-todo/internal/storage"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Не хватает аргументов")
		os.Exit(0)
	}

	var tasks = []models.Task{}
	storage.Load(&tasks)
	defer storage.Save(&tasks)

	fmt.Println(tasks)
}
