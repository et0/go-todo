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

		now := time.Now().Local()
		tasks = append(tasks, models.Task{
			Id:        len(tasks) + 1,
			Title:     os.Args[2],
			Completed: 0,
			CreatedAt: now.Format("02.01.2006"),
		})
	case "list":
		isCompleted := false
		isCheck := " "
		if len(os.Args) == 3 && os.Args[2] == "--completed" {
			isCompleted = true
			isCheck = "\u2713"
		}
		for _, t := range tasks {
			if (isCompleted && t.Completed == 0) || (!isCompleted && t.Completed == 1) {
				continue
			}
			fmt.Printf("%d. [%s] %s (%s)\n", t.Id, isCheck, t.Title, t.CreatedAt)
		}
	case "done":
	case "delete":
	case "clear":
	}
}
