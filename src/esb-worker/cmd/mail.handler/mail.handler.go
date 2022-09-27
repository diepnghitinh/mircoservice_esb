package mail_handler

import (
	"os"
	"time"

	// "strconv"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	gomail "gopkg.in/mail.v2"

	"esb-worker/cmd/mail.handler/env"
	schema "esb-worker/cmd/mail.handler/schemas"
	"esb-worker/internal/core/libs"
)

var log = logrus.New()

type Worker struct {
	Http *resty.Request
	Env  *env.Config
}

func (c *Worker) Init(esb *libs.ESB) {
	c.Env = &env.Config{}
	_ = esb.LoadConfig(c.Env)
	c.Http = resty.New().SetRetryCount(3).R().EnableTrace()
}

func (c *Worker) Handler(d amqp.Delivery) error {
	// Parse schema
	jsonData := &schema.Message{}
	time.Sleep(10 * time.Second)

	if err := json.Unmarshal(d.Body, &jsonData); err != nil {
		log.Fatal(err)
		return err
	}
	// Send email
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", c.Env.Services.Username)
	m.SetAddressHeader(
		"From", c.Env.Services.Username, c.Env.Services.SenderName)

	// Set E-Mail receivers
	m.SetHeader("To", jsonData.To)

	// Set E-Mail subject
	m.SetHeader("Subject", jsonData.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", jsonData.Body)

	// Set Attachments
	dir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(dir) // clean up

	for i := range jsonData.Attachments {
		b64File := jsonData.Attachments[i]
		fileContent, err := base64.StdEncoding.DecodeString(
			b64File.FileContent)

		if err != nil {
			log.Fatal(err)
			return err
		}

		tempFile := filepath.Join(dir, b64File.FileName)
		ioutil.WriteFile(tempFile, fileContent, 0644)
		m.Attach(tempFile)
	}

	// Settings for SMTP server
	smtpDial := gomail.NewDialer(
		c.Env.Services.Host,
		c.Env.Services.Port,
		c.Env.Services.Username,
		c.Env.Services.Password)

	// Now send E-Mail
	if err := smtpDial.DialAndSend(m); err != nil {
		return err
	}

	return nil

}

func Init() interface{} {
	var worker libs.IConsumer
	worker = &Worker{}
	return worker
}
