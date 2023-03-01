package main

import (
	//"fmt"
	"net/http"
	//"html/template"
	//"log"
	//"encoding/json"
	"path"
	"path/filepath"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

	//Configure Gin server object
	r := gin.Default()

	//Use cors package to enable cross-origin resource sharing

	config := cors.DefaultConfig()
 
 	config.AllowHeaders = []string{"Authorization", "content-type"}
 	config.AllowOrigins = []string{"http://localhost:4200"}

	//Since Gin requires specific paths, ensures Angular files are served with specific roots
    r.NoRoute(func(c *gin.Context) {
        dir, file := path.Split(c.Request.RequestURI)
        ext := filepath.Ext(file)
        if file == "" || ext == "" {
            c.File("./ui/dist/ui/index.html")
        } else {
            c.File("./ui/dist/ui/" + path.Join(dir, file))
        }
    })

	/**------------------------------------------------------------------------
	 *                           AUTHENTICATION
	 *------------------------------------------------------------------------**/

	//Redirect to third-party authentication service
	r.GET("/auth/{provider}", func(c *gin.Context) {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})

	//Validate login with third-party
	r.GET("/callback", func(c *gin.Context) {
		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Authentication successful",
			"id": user.UserID,
			"email" : user.Email,
			"name" : user.Name,
		})
	})

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
	/*---------------------------- END OF AUTH ROUTING ----------------------------*/
}