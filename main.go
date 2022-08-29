package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BobRusch/GoAuthService/handlers"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	os.Setenv("SERVER_PORT", os.Getenv("SERVER_PORT"))
	os.Setenv("GOOGLE_OAUTH_CLIENT_ID", os.Getenv("GOOGLE_OAUTH_CLIENT_ID"))
	os.Setenv("GOOGLE_OAUTH_CLIENT_SECRET", os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"))
	os.Setenv("FACEBOOK_OAUTH_CLIENT_ID", os.Getenv("FACEBOOK_OAUTH_CLIENT_ID"))
	os.Setenv("FACEBOOK_OAUTH_CLIENT_SECRET", os.Getenv("FACEBOOK_OAUTH_CLIENT_SECRET"))
	os.Setenv("FACEBOOK_REDIRECT_URL", os.Getenv("FACEBOOK_REDIRECT_URL"))

}
func main() {

	InitEnv()

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", svr.Addr)
	if err := svr.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
