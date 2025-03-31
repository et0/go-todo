package main

import (
	"fmt"
	"os"
	"strconv"
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
		lastId := tasks[len(tasks)-1].Id
		tasks = append(tasks, models.Task{
			Id:        lastId + 1,
			Title:     os.Args[2],
			Completed: 0,
			CreatedAt: now.Format("02.01.2006"),
		})
		fmt.Printf(models.SuccessAdd, os.Args[2])
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
		if len(os.Args) < 3 {
			fmt.Println(models.WrongArgsDone)
			os.Exit(0)
		}
		doneId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(models.WrongArgsDone)
			os.Exit(0)
		}
		for i, t := range tasks {
			if t.Id != doneId {
				continue
			}
			tasks[i].Completed = 1
			break
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println(models.WrongArgsDelete)
			os.Exit(0)
		}
		doneId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(models.WrongArgsDelete)
			os.Exit(0)
		}
		for i, t := range tasks {
			if t.Id != doneId {
				continue
			}

			// Если это последний элемент
			if i+1 == len(tasks) {
				tasks = tasks[:i]
			} else {
				tasks = append(tasks[:i], tasks[i+1:]...)
			}
			fmt.Printf(models.SuccessDelete, t.Title)
			break
		}
	case "clear":
	}
}
