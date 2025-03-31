package models

const (
	WrongArgs       = "Не хватает аргументов"
	WrongArgsAdd    = "Использовать команду: todo add <TITLE>"
	WrongArgsDone   = "Использовать команду: todo done <ID>"
	WrongArgsDelete = "Использовать команду: todo delete <ID>"

	SuccessAdd    = "\u2705 Новая задача: \"%s\" добавлена\n"
	SuccessDelete = "\u274C Задача: \"%s\" удалена\n"
	SuccessClear  = "\u274C Удалено: %d задач\n"
)
