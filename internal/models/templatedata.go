package models

import "github.com/aleph-design/startbridge/internal/forms"

// TemplateData holds data sent from handlers to templates
// It holds information available to every single page in
// the whole application.
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Compy     map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	DispImg   bool
	Form      *forms.Form
	IsAuth    bool
}
