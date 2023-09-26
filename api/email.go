package api

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

type Body struct {
	To string `json:"email_to"`
}

func SendEmail(c *gin.Context) {
	var data Body
	if err := c.BindJSON(&data); err != nil {
		//Handle error
	}

	if err := godotenv.Load(".env"); err != nil {
		//Handle error
	}
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailTo := data.To
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Generate a random Message-ID
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	messageID := strconv.FormatInt(r.Int63(), 10) + "@" + smtpHost

	messageBody := "Thank you for reading this article."

	message := "From: " + emailFrom + "\n" +
		"To: " + emailTo + "\n" +
		"Subject: " + "This is a subject" + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n" +
		"Message-ID: <" + messageID + ">\n\n" +
		messageBody

	// Set up authentication
	auth := smtp.PlainAuth("", emailFrom, emailPassword, smtpHost)

	// Without TLS
	/* err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailFrom, []string{emailTo}, []byte(message))
	if err != nil {
		//Handle error
	} */


	/* Using TLS */
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		//Handle error
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		//Handle error
	}

	// Authenticate
	if err := client.Auth(auth); err != nil {
		//Handle error
	}

	// Set the sender and recipient
	if err := client.Mail(emailFrom); err != nil {
		//Handle error
	}
	if err := client.Rcpt(emailTo); err != nil {
		return
	}

	// Send the email body
	wc, err := client.Data()
	if err != nil {
		//Handle error
	}

	_, err = wc.Write([]byte(message))
	if err != nil {
		//Handle error
	}
	err = wc.Close()
	if err != nil {
		//Handle error
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"response": "success",
		"message":  "email sent",
	})

}
