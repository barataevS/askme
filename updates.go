package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	token  = "5963202075:AAFrnFHCrvZ-2vu8ieKI-eU3D6PX8IpDcpw"
	urlAPI = "https://api.telegram.org/bot"
)

func Updates(url, offset string) ([]Update, error) {
	request, err := http.Get(url + offset)
	if err != nil {
		log.Printf("Ошибка с получением обновления от Телеграм")
	}

	defer request.Body.Close()
	var data Data

	if err = json.NewDecoder(request.Body).Decode(&data); err != nil {
		log.Printf("Ошибка с расшифровкой обновления от Телеграм")
	}

	return data.Data, nil
}

func Respond(url string, update Update) error {
	x := getAnswerFromChat(update.Message.Text)
	var responseTier1 Response
	responseTier1.Id = update.Message.Chat.Id
	responseTier1.RespondMessage = x
	responseTier1.Links = true
	//responseTier1.RespondMessage, responseTier1.ReplyMarkup.Keyboard, responseTier1.ReplyMarkup.OneTimeKey,
	//	responseTier1.Parse = ResponseTierOne(update, responseTier1)

	buf, err := json.Marshal(responseTier1)
	if err != nil {
		log.Printf("Ошибка с получением обновления от Телеграм")
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Printf("Ошибка с получением обновления от Телеграм")
	}
	fmt.Println("Ответ направлен, ", update.Message.Chat.Name, update.Message.Chat.Id, ", ", time.Now())
	return nil

}
