package main

import (
	"log"
)

func createDatabase(username string, email string, userID string) {
	log.Println("Authorization successful!")
	log.Println("Username: ", username)
	log.Println("Email: ", email)
	log.Println("UID: ", userID)
}