package main

import (
	"log"

	"github.com/dayroromero/stori-challenge/pkg/notifications"
	"github.com/dayroromero/stori-challenge/utils"
)

func main() {

	utils.LoadEnv()
	//csv.File_processor()
	var email notifications.EmailNotification
	email.RecipientEmail = "dayro.romero@outlook.com"
	email.Subject = "Transactions Summary"
	err := notifications.OrchestrateEmailSending(email)
	if err != nil {
		log.Printf("Error email ochestation: %v", err)
	}
}
