package main

import (
 "os"
 "text/template"
)

// Структура для создания шаблона
type Todo struct {
 Task string
}

func main() {
 // Шаблон для генерации данных
 tmpl, err := template.New("todos").Parse("Task: {{.Task}}\n")
 if err != nil {
  panic(err)
 }

 // Список задач
 todos := []Todo{
  {"Task 1"},
  {"Task 2"},
  {"Task 3"},
 }

 // Применяем шаблон к каждой задаче
 for _, todo := range todos {
  err := tmpl.Execute(os.Stdout, todo)
  if err != nil {
   panic(err)
  }
 }
}
