package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/aleph-design/startbridge/internal/config"
)

// site wide access to app config
var app *config.AppConfig

// actually populate this package with site-wide
// access to AppConfig by transferring a pointer
// *config.AppConfig and call this function from
// main()
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	// Demonstrating that we can now write to a log
	app.InfoLog.Println("Client error with status: ", status)
	http.Error(w, http.StatusText(status), status)
}

// In production we'll the log write to a file and
// put a directive that send a text or an email to
// the administrator of the site
func ServerError(w http.ResponseWriter, err error) {
	// We can print a trace of the errror
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// For now just write to the error log
	app.ErrorLog.Println(trace)
	// We need to give the user some kind of feedback
	http.Error(w, http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}
