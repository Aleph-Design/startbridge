package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Aleph-Design/startbridge/internal/driver"
	"github.com/aleph-design/startbridge/internal/bidgame"
	"github.com/aleph-design/startbridge/internal/config"
	"github.com/aleph-design/startbridge/internal/controller"
	"github.com/aleph-design/startbridge/internal/handlers"
	"github.com/aleph-design/startbridge/internal/helpers"
	"github.com/aleph-design/startbridge/internal/models"
	"github.com/aleph-design/startbridge/internal/render"
	"github.com/alexedwards/scs/v2"
)

const (
	MAILFROM   = "me@here.nl"
	DOMAINNAME = "http://localhost:5000"
	PORTNUMBER = ":5000"
	MAILCRYPT  = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
	//  01234567890123456789012345678901
	SUBJECT = "Registratie Startbridge"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {
	// here we tell the engine which types is sent that implments
	// the interfaces we store in session
	// gob.Register(models.Home{})
	// gob.Register(models.People{})
	// gob.Register(models.Company{})

	// seed the randomizer only ones.
	// rand.Seed(time.Now().UnixNano())

	/*********************************
	* CHANGE THIS WHEN IN PRODUCTION *
	*********************************/
	app.InProduction = false

	// mailData := models.MailData{
	// 	From:       MAILFROM,
	// 	Subject:    SUBJECT,
	// 	DomainName: DOMAINNAME,
	// 	MailCrypt:  MAILCRYPT,
	// }
	// app.MailData = &mailData

	// SETUP SOME LOGGERS:
	// infolog prints to the standard out: the terminal
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	// errorlog for now also prints to the terminal
	// ********** CHANGE THAT IN PRODUCTION ******************
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	// also see internal/helpers/helpers.go

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// add session to AppConfig
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// create a site wide pointer to BidData
	app.BidData = &models.BidData{}

	// initialize database repo & connect to database
	log.Println("connecting to database . . .")
	connStr := "host=localhost port=5432 dbname=startbridge user=janhkila password="

	db, err := driver.ConnectSQL(connStr)
	// db, err := driver.ConnectSQL(connStr)
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	defer db.SQL.Close()
	log.Println("connected to database")

	// create channel to send email
	// mailChan := make(chan models.MailData)
	// app.MailChan = mailChan
	// defer close(app.MailChan)
	// listenForMail()
	// fmt.Println("Mail Listener started . . .")

	app.TemplateCache = tc
	app.UseCache = false

	// repo := handlers.NewRepo(&app, db)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// give these packages access to *config.AppConfig
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)
	bidgame.NewBidGame(&app)
	controller.NewController(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORTNUMBER))

	srv := &http.Server{
		Addr:    PORTNUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
