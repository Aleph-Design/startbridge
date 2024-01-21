package main

import (
	"net/http"

	"github.com/aleph-design/startbridge/internal/config"
	"github.com/aleph-design/startbridge/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// ignore any request - that is a post - and does not
	// have a proper Cross Site Request Token.
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)

	mux.Get("/newbidgame", handlers.Repo.NewBidGame)
	mux.Get("/bidbox", handlers.Repo.Bidbox)

	mux.Get("/over", handlers.Repo.Over)

	mux.Get("/user-login", handlers.Repo.UserLogin)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/registration", handlers.Repo.Registration)
	mux.Post("/registration", handlers.Repo.PostRegistration)

	// mux.Post("/people-login", handlers.Repo.PostShowLogin)
	// mux.Get("/people-logout", handlers.Repo.Logout)



	// tell GO where to find static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
