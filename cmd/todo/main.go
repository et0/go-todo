package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/et0/go-todo/internal/models"
	"github.com/et0/go-todo/internal/storage"

	"github.com/jackc/pgx/v5"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(models.WrongArgs)
		os.Exit(0)
	}

	conn, err := pgx.Connect(context.Background(), "postgresql://leonid:@localhost:5432/todo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}

	var tasks = []models.Task{}
	storage.LoadTasks(conn, &tasks)

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println(models.WrongArgsAdd)
			os.Exit(0)
		}

		err := storage.InsertTask(conn, os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

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
		for _, t := range tasks {
			if t.Id != doneId {
				continue
			}

			err := storage.SetDone(conn, t.Id)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
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
		err := storage.DeleteAll(conn)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
