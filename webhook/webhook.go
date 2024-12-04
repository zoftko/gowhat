package webhook

type Notification struct {
	Object string   `json:"object"`
	Entry  []Change `json:"entry"`
}

type Change struct {
	ID      string        `json:"id"`
	Changes []ValueObject `json:"changes"`
}

type ValueObject struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string    `json:"messaging_product"`
	Metadata         Metadata  `json:"metadata"`
	Contacts         []Contact `json:"contacts"`
	Messages         []Message `json:"messages"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Contact struct {
	WhatsappID string  `json:"wa_id"`
	Profile    Profile `json:"profile"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Text      Text   `json:"text"`
}

type Text struct {
	Body string `json:"body"`
}
