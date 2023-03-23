package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/garrettjhnsn/projectmf/helpers"
	"github.com/garrettjhnsn/projectmf/internal/config"
	"github.com/garrettjhnsn/projectmf/internal/driver"
	"github.com/garrettjhnsn/projectmf/internal/handlers"
	"github.com/garrettjhnsn/projectmf/internal/models"
	"github.com/garrettjhnsn/projectmf/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	log.Println("Starting Web Application on Port:", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	//Session
	gob.Register(models.ConsultationRequest{})

	//Change To True In Production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//Connect To Database
	log.Println("Connecting to Database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=projectmf user=postgres password=Noodle95$$")
	if err != nil {
		log.Fatal("Cannot Connect to Database")
	}
	log.Println("Successfully Connected to Database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}
	app.TemplateCache = tc

	// False = DevMode On [Cache off] True = DevMode Off [Cache On]
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
