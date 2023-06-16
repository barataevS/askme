package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getToken(token string) string {
	return urlAPI + token + "/getUpdates"
}

func postToken(token string) string {
	return urlAPI + token + "/sendMessage"
}

func getOffset(offset int) string {
	return "?offset=" + strconv.Itoa(offset)
}

func main() {
	// Задайте текст запроса
	//	input := "Столица России?"
	var offset int

	// Создайте структуру для передачи запроса в JSON формате

	for {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Ошибка 1")
			}
		}()
		data, err := Updates(getToken(token), getOffset(offset))

		if err != nil {
			log.Printf("Ошибка 2")
		}
		//fmt.Println(moexNew.ResponseSlice)

		for _, y := range data {
			fmt.Println(y.Message)
			//go sendChannel(buf1, y.Message.Text)
			//time.Sleep(100 * time.Millisecond)
			err = Respond(postToken(token), y)

			offset = y.Id + 1

		}
	}
}

func getAnswerFromChat(vopros string) string {

	requestOne := Requests{
		Model: "gpt-3.5-turbo",
		Messages: []MessageOne{
			{
				Role:    "user",
				Content: vopros,
			},
		},
	}

	// Преобразуйте структуру запроса в JSON
	jsonData, err := json.Marshal(requestOne)
	if err != nil {
		fmt.Println("Ошибка при маршалинге запроса:", err)
		return ""
	}

	// Определите адрес API и заголовки
	apiURL := "https://api.openai.com/v1/chat/completions"
	apiKey := "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"

	// Создайте HTTP POST запрос к API
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return ""
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Отправьте запрос к API
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return ""
	}
	defer response.Body.Close()

	// Прочитайте ответ от API
	//responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return ""
	}
	var jsonResponse GPTResponse
	// Прочитайте ответ от API
	if err = json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("Ошибка с раскодированием ответа")
	}

	// Выведите текст ответа
	x := jsonResponse.Choices[0].Message.Content

	fmt.Println("Ответ от ChatGPT:", x)
	return x
}
