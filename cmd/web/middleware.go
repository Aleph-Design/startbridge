package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// This sets a cookie that's available on a per page basis.
	// So, the browser sets a different cookie for every page.
	// See: https://www.alexedwards.net/blog/working-with-cookies-in-go
	// http.Cookie represents in GO: type Cookie struct {...}
	csrfHandler.SetBaseCookie(http.Cookie {
		// Name and Value are mandatory; they're set by NoSurf package
		HttpOnly: true,
		Path:     "/",									// apply to the entire site
		Secure:   app.InProduction,			// set in main.go when in production
		SameSite: http.SameSiteLaxMode,	// value from net/http package
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// Our own peace of custom middleware that has access to the request
// so that we can call helpers.
// we protect routes in the routes file to insure that only visitors
// taht are logged in have access to routes that we want to protect.
// func Auth(next http.Handler) http.Handler {
// 	// here we need access to the request because we call the
// 	// helper function IsAuth()
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// now we can call the helper function
// 		if !helpers.IsAuth(r) {
// 			session.Put(r.Context(), "error", "Eerst inloggen!")
// 			// redirect to the login screen
// 			http.Redirect(w, r, "/people-login", http.StatusSeeOther)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
//}
