package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rgomezs4/event_registration/controllers"
	"github.com/rgomezs4/event_registration/data"

	"github.com/gorilla/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3158"
	}

	dn := os.Getenv("APP_DB_DRIVER")
	if dn == "" {
		dn = "postgres"
	}

	ds := os.Getenv("APP_DB_SOURCE")

	api := controllers.NewAPI()
	// open the database connection
	db := &data.DB{}

	if err := db.Open(dn, ds); err != nil {
		log.Fatal("unable to connect to the database:", err)
	}

	api.DB = db

	log.Println("Server started on port " + port)

	if err := http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "application-token", "AUTH_USER_ID", "X_FILE_NAME"}),
		handlers.ExposedHeaders([]string{"Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(api)); err != nil {
		log.Println(err)
	}
}
