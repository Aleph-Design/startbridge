package config

import (
	"html/template"
	"log"

	// Cannot point to bidmap because of import cycle!

	"github.com/aleph-design/startbridge/internal/bidmap"
	"github.com/aleph-design/startbridge/internal/models"
	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	Game          *models.Game
	BidIndex      *bidmap.BidIndex
	BidData       *models.BidData
	// MailChan      chan models.MailData
	// MailMsg       models.MailData
	// MailData      *models.MailData
}
