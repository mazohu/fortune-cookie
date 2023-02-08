package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"

	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

func main() {
	//Read in the environment variables from .env for security
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	key := "SESSIONS_SECRET"
	duration := 86400 * 30  // Set session expiry date to 30 days
	
	//Configure session
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(duration)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true //True when serving over https

	gothic.Store = store

	goth.UseProviders(google.New(viper.GetString("GOOGLE_CLIENT_ID"), viper.GetString("GOOGLE_CLIENT_SECRET"), "http://localhost:3000/auth/google/callback", "email", "profile"))

	p := pat.New()
	//Redirect to Google sign-in
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
		fmt.Fprintln(res, err)
		return
		}
		log.Println("Authorization successful!")
		log.Println("Username: ", user.Name)
		log.Println("Email: ", user.Email)
		log.Println("UID: ", user.UserID)

	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("../index.html")
		t.Execute(res, false)
	})
	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}
