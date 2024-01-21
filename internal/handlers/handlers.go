package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aleph-design/startbridge/internal/bidmap"
	"github.com/aleph-design/startbridge/internal/config"
	"github.com/aleph-design/startbridge/internal/forms"
	"github.com/aleph-design/startbridge/internal/models"
	"github.com/aleph-design/startbridge/internal/render"
	"github.com/aleph-design/startbridge/internal/repository"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// // NewRepo creates a new repository
//
//	func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
//		return &Repository{
//			App: a,
//			DB:  dbrepo.NewPostgresRepo(db.SQL, a),
//		}
//	}
//
// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		// DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

/************************
*						HOME
*************************/
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// we want to re-display already inputted data
	// var emptyHome models.Home
	// data := make(map[string]interface{})
	// data["home"] = emptyHome

	render.Template(w, r, "home.page.tmpl", &models.TemplateData{
		// include an empty Form object the first time this page is rendered
		Form:    forms.New(nil),
		DispImg: true, // displays top picture with text
	})
}

func (m *Repository) Over(w http.ResponseWriter, r *http.Request) {
	// we want to re-display already inputted data
	// var emptyHome models.Home
	// data := make(map[string]interface{})
	// data["home"] = emptyHome

	render.Template(w, r, "over.page.tmpl", &models.TemplateData{
		// include an empty Form object the first time this page is rendered
		Form:    forms.New(nil),
		DispImg: true, // displays top picture with text
	})
}

/*
***************************

	START A NEW BIDDING PHASE

****************************
*/
func (m *Repository) NewBidGame(w http.ResponseWriter, r *http.Request) {

	// Initiate data and processing structures
	// =======================================

	// TODO => block page refresh *************************************

	// create a new game model every time we begin a new bidding round
	// the old one is taken care of by GC
	// game := models.Game{}
	// g := &game
	// m.App.Game = g

	// _ = bidgame.InitNewGame(g)
	// // bidgame.ShowSouthHand(g)

	// // fmt.Println("\nhandlers - start a new bidding")
	// m.App.BidIndex = &bidmap.BidIndex{}
	// m.App.BidIndex.Count = 0

	// m.App.BidData = &models.BidData{} // BidData is now empty
	// data := make(map[string]interface{})
	// m.App.BidData.TopMsg = "Wat bied je?"
	// m.App.BidData.Msg_1 = "Het eerste bod in deze biedronde."
	// data["bidData"] = *m.App.BidData

	// This is the first time newbidgame.page.tmpl is rendered
	// It should be the only time during the bidding process.
	render.Template(w, r, "newbidgame.page.tmpl", &models.TemplateData{
		// Data:    data,
		DispImg: false, // removes top picture and text
	})
}

// repeatedly called after each click in the bidbox
// ================================================
func (m *Repository) Bidbox(w http.ResponseWriter, r *http.Request) {
	// now, there's a click in the bidbox that we want to get
	id := r.URL.Query().Get("id")
	bidID, _ := strconv.Atoi(id) // var bidID is created
	// fmt.Println("\nbidID: ", bidID)

	bid := bidmap.NewBid(bidID)
	fmt.Println("bid: ", bid)

	data := make(map[string]interface{})
	// BidDate is set in helpers check valid bid
	data["bidData"] = *m.App.BidData
	render.Template(w, r, "newbidgame.page.tmpl", &models.TemplateData{
		Data:    data,
		DispImg: false, // removes top picture and text
	})
	// http.Redirect(w, r, "/", http.StatusSeeOther)
} // BidBox() =============================================

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// we want to re-display already inputted data
	// var emptyHome models.Home
	// data := make(map[string]interface{})
	// data["home"] = emptyHome

	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{
		// include an empty Form object the first time this page is rendered
		Form:    forms.New(nil),
		DispImg: true, // displays top picture with text
	})
}

/****************************
*  		USER REGISTRATION
*****************************/
func (m *Repository) Registration(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "registration.page.tmpl", &models.TemplateData{
		Form:    forms.New(nil), // we need the empty form!
		DispImg: false,          // removes top picture and text
	})
}

func (m *Repository) PostRegistration(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// helpers.ServerError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/***************************
*     SHOW LOGIN / LOGOUT
****************************/
func (m *Repository) UserLogin(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "user-login.page.tmpl", &models.TemplateData{
		Form:    forms.New(nil), // we need the empty form!
		DispImg: true,
	})
}
