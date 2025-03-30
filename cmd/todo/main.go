package main

import (
	"fmt"
	"os"
	"time"

	"github.com/et0/go-todo/internal/models"
	"github.com/et0/go-todo/internal/storage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(models.WrongArgs)
		os.Exit(0)
	}

	var tasks = []models.Task{}
	storage.Load(&tasks)
	defer storage.Save(&tasks)

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println(models.WrongArgsAdd)
			os.Exit(0)
		}
		now := time.Now()
		tasks = append(tasks, models.Task{Id: len(tasks) + 1, Title: os.Args[2], Complated: 0, CreatedAt: int(now.Unix())})
	case "list":
	case "done":
	case "delete":
	case "clear":
	}

	fmt.Println(tasks)
}
