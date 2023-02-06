package main

import (
	"os"

	"github.com/markbates/goth/providers/google"
)

func main() {
	google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:3000/auth/google/callback")
}
