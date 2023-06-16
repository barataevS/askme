package main

type GPTResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

//	type Choice struct {
//		Message struct {
//			//Role    string `json:"role"`
//			Content string `json:"content"`
//		} `json:"message"`
//	}
type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content string `json:"content"`
}

type Request struct {
	Prompt string `json:"prompt"`
}

type MessageOne struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Requests struct {
	Model    string       `json:"model"`
	Messages []MessageOne `json:"messages"`
}

type Update struct {
	Id      int      `json:"update_id"`
	Message Messages `json:"message"`
}

type Messages struct {
	Chat      Chat   `json:"chat"`
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
}

type Chat struct {
	Id   int    `json:"id"`
	Name string `json:"first_name"`
}

type Data struct {
	Data []Update `json:"result"`
}

type Response struct {
	Id             int    `json:"chat_id"`
	RespondMessage string `json:"text"`
	//ReplyMarkup    ReplyKeyboardMarkup `json:"reply_markup"`
	Parse string `json:"parse_mode"`
	Links bool   `json:"disable_web_page_preview"`
}
type ReplyKeyboardMarkup struct {
	Keyboard   [][]KeyboardButton `json:"keyboard"`
	OneTimeKey bool               `json:"resize_keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type ItemTass struct {
	Date     string `xml:"pubDate"`
	Category string `xml:"category"`
	Title    string `xml:"title"`
	Link     string `xml:"link"`
}
