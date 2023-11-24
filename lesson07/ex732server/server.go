package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчики для разных путей
	http.HandleFunc("/", handleRequest)

	// Запускаем веб-сервер на порту 9999
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// В зависимости от метода HTTP-запроса вызываем соответствующий обработчик
	switch r.Method {
	case http.MethodGet:
		handleGET(w, r)
	case http.MethodPost:
		handlePOST(w, r)
	case http.MethodPut:
		handlePUT(w, r)
	case http.MethodDelete:
		handleDELETE(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
	fmt.Println(r.Method) // Тип метода
	fmt.Println(r.URL)    // запрашиваемый URL
	fmt.Println(r.Proto)  // Протокол
}

// Обработчик для GET-запросов
func handleGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это GET-запрос!")
}

// Обработчик для POST-запросов
func handlePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это POST-запрос!")
}

// Обработчик для PUT-запросов
func handlePUT(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это PUT-запрос!")
}

// Обработчик для DELETE-запросов
func handleDELETE(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это DELETE-запрос!")
}