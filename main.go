//* Below is me just trying to get personally made packages to work ignore it
// package main

// import (
// 	"fmt"
// 	//"ourProject/helper"
// )

// func main() {
//     fmt.Println("hello world")
// 	//helper.Help();
// }

//* Below is the start of the tutorial

package main

import (

	"fmt"
	"os"
	"log"
  	"github.com/joho/godotenv"
)


// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func envVariable(key string) string {

	// set env variable using os package
	os.Setenv(key, "gopher")
  
	// return the env variable using os package
	return os.Getenv(key)
  }

func main() {

	// os package
	value := envVariable("name")
  
	fmt.Printf("os package: name = %s \n", value)
	fmt.Printf("environment = %s \n", os.Getenv("APP_ENV"))


  // godotenv package
  dotenv := goDotEnvVariable("STRONGEST_AVENGER")

  fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)
}

// package main

// import (
//   "fmt"
//   "html/template"
//   "net/http"
  

//   "log"

//   "github.com/gorilla/pat"
//   "github.com/markbates/goth"
//   "github.com/markbates/goth/gothic"
//   "github.com/markbates/goth/providers/google"
//   "github.com/gorilla/sessions"
//   "github.com/joho/godotenv"
//   "os"
// )

// func goDotEnvVariable(key string) string {

// 	// load .env file
// 	err := godotenv.Load(".env")
  
// 	if err != nil {
// 	  log.Fatalf("Error loading .env file")
// 	}
  
// 	return os.Getenv(key)
//   }


// func main() {
  
//   key := goDotEnvVariable("SECRET-SESSION-KEY")  // Replace with your SESSION_SECRET or similar
//   maxAge := 86400 * 30  // 30 days
//   isProd := false       // Set to true when serving over https

//   store := sessions.NewCookieStore([]byte(key))
//   store.MaxAge(maxAge)
//   store.Options.Path = "/"
//   store.Options.HttpOnly = true   // HttpOnly should always be enabled
//   store.Options.Secure = isProd

//   gothic.Store = store

//   goth.UseProviders(
//     google.New("our-google-client-id", "our-google-client-secret", "http://localhost:3000/auth/google/callback", "email", "profile"),
//   )

//   p := pat.New()
//   p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

//     user, err := gothic.CompleteUserAuth(res, req)
//     if err != nil {
//       fmt.Fprintln(res, err)
//       return
//     }
//     t, _ := template.ParseFiles("templates/success.html")
//     t.Execute(res, user)
//   })

//   p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
//     gothic.BeginAuthHandler(res, req)
//   })

//   p.Get("/", func(res http.ResponseWriter, req *http.Request) {
//     t, _ := template.ParseFiles("templates/index.html")
//     t.Execute(res, false)
//   })
//   log.Println("listening on localhost:3000")
//   log.Fatal(http.ListenAndServe(":3000", p))
// }