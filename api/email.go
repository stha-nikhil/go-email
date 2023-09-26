package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"math/rand"
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

}
