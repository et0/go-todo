package models

type Task struct {
	Id        int    `json:"id"`         // Идентификатор задачи
	Title     string `json:"title"`      // Описание задачи
	Complated int    `json:"complated"`  // Статус выполнения 0 или 1
	CreatedAt string `json:"created_at"` // Время добавления задачи
}
