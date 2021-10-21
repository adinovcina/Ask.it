package service

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(purpose string, title string, email string, content string, username string) {
	from := mail.NewEmail(purpose, "ovcinaadin@gmail.com")
	to := mail.NewEmail("User to", email)
	htmlContent := fmt.Sprintf("<h2>Hey!, new registration on your account noticed</h2>"+
		"<br/><h3>Hello <span style='color:blue'>%s</span>, your password is : "+
		"<span style='color:blue'>%s</span></h3><br/><h5>Registered at : %s</h5>", content, username, time.Now().Format("2006-01-02 15:04:05"))
	message := mail.NewSingleEmail(from, title, to, content, htmlContent)
	sDec, _ := b64.StdEncoding.DecodeString(os.Getenv("API_KEY"))
	apiKey := string(sDec)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
