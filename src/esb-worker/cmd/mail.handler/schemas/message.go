package schema

type Attachments struct {
	FileName string `json:"filename`
	FileContent  string `json:"filecontent"`
}

type Message struct {
	To          string        `json:"to"`
	Subject     string        `json:"subject"`
	Body        string        `json:"body"`
	BodyType    string        `json:"body_type"`
	Attachments []Attachments `json:"attachments`
}
