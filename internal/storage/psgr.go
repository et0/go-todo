package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/et0/go-todo/internal/models"
	"github.com/jackc/pgx/v5"
)

func InsertTask(conn *pgx.Conn, title string) error {
	now := time.Now().Local()
	_, err := conn.Exec(context.Background(), "INSERT INTO tasks (title, completed, created_at) VALUES ($1, 0, $2)",
		title, now.Format("02.01.2006"))
	if err != nil {
		return fmt.Errorf("[INSERT] %w", err)
	}

	return nil
}

func SetDone(conn *pgx.Conn, id int) error {
	_, err := conn.Exec(context.Background(), "UPDATE tasks SET completed = 1 WHERE id = $1",
		id)
	if err != nil {
		return fmt.Errorf("[SET COMPLETE] %w", err)
	}

	return nil
}

func DeleteAll(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM tasks")
	if err != nil {
		return fmt.Errorf("[DELETE] %w", err)
	}

	return nil
}

func LoadTasks(conn *pgx.Conn, t *[]models.Task) error {
	rows, err := conn.Query(context.Background(), "SELECT * FROM tasks ORDER BY id DESC")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Completed,
			&task.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("error scanning task row: %w", err)
		}
		*t = append(*t, task)
	}

	return nil
}
