package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// User - структура для представления объекта пользователя
type User struct {
	Name string `json:"name"`
	ID   uint32 `json:"id"`
}

func main() {
	// Создаем экземпляр структуры User
	var u = User{
		Name: "Alex",
		ID:   1,
	}

	// Кодируем структуру User в JSON (байтовый срез)
	bytesRepresentation, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	// Отправляем POST-запрос на сервер с JSON-телом
	// resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	resp, err := http.Post("http://localhost:9999/body", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	// Читаем и конвертируем тело ответа в байты
	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Выводим содержимое тела ответа
	fmt.Println(string(bytesResp))
}
//curl -X POST "http://localhost:8080/body" -H "accept: application/json" -d '{"login":"my_login","password":"my_password"}'