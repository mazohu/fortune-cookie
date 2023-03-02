package main

import (
	"net/http"
	"log"
	//"encoding/json"
	//"path"
	//"path/filepath"

	"fmt"
	"html/template"
	"github.com/gorilla/pat"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	
	// "github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
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

	goth.UseProviders(google.New(viper.GetString("GOOGLE_CLIENT_ID"), viper.GetString("GOOGLE_CLIENT_SECRET"), "http://localhost:4200/auth/google/callback", "email", "profile"))

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
		fmt.Fprintln(res, err)
		return
		}
		t, _ := template.ParseFiles("templates/success.html")
		t.Execute(res, user)

		//~ Start of testing database functions
		//accessDatabase(user.Name, user.Email, user.UserID)
		//dataBaseTesting(user.Name, user.Email, user.UserID)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("./index.html")
		t.Execute(res, false)
	})
	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))

	//Configure Gin server object
	// r := gin.Default()

	// //Use cors package to enable cross-origin resource sharing

	// config := cors.DefaultConfig()
 
 	// config.AllowHeaders = []string{"Authorization", "content-type"}
	// //Client is open on 4200
 	// config.AllowOrigins = []string{"http://localhost:4200"}

	// /**------------------------------------------------------------------------
	//  *                           AUTHENTICATION
	//  *------------------------------------------------------------------------**/

	// //!Something with routing is severely wrong

	// //Redirect to third-party authentication service
	// r.GET("/auth/{provider}", func(c *gin.Context) {
	// 	gothic.BeginAuthHandler(c.Writer, c.Request)
	// 	log.Println("goth auth began")
	// })

	// //Validate login with third-party
	// r.GET("/callback", func(c *gin.Context) {
	// 	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	// 	if err != nil {
	// 		c.AbortWithError(http.StatusInternalServerError, err)
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Authentication successful",
	// 		"id": user.UserID,
	// 		"email" : user.Email,
	// 		"name" : user.Name,
	// 	})
	// })

	// //Server listens on port 3000 => client-side requests should be made to 3000
	// err := r.Run(":3000")
	// if err != nil {
	// 	panic(err)
	// }
	// /*---------------------------- END OF AUTH ROUTING ----------------------------*/
}