package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	data := map[string]string{
		"Email":    "dimon2020@mail.ru",
		"Password": "123456",
	}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(
		"http://localhost:8080/auth",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
