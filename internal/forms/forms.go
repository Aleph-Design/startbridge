package forms

import (
	"net/http"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

// holds all the forms information.
// Don't forget to add member "Form *forms.Form"
// to TemplateData - this way we know that the
// Form object is available on every page.
type Form struct {
	url.Values
	Errors errors
}

// returns true when there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// initializes a new Form struct; it will be empty the
// first time the page is rendered.
func New(data url.Values) *Form {
	return &Form{
		// corresponding to the inputs on the form field
		data,
		errors(map[string][]string{}),
	}
}

// create a variatic function, we can put as many strings
// as we want is parameters.
func (f *Form) Required(fields ...string) {
	// we range trough the fields variable
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Mag niet leeg zijn!")
		}
	}
}

// check both fields empty
func (f *Form) BothEmpty(field1, field2 string) bool {
	valOne := strings.TrimSpace(f.Get(field1))
	valTwo := strings.TrimSpace(f.Get(field2))
	if valOne == "" && valTwo == "" {
		f.Errors.Add(field1, "U moet ÉÉN der velden vullen!")
		f.Errors.Add(field2, "U moet ÉÉN der velden vullen!")
		return false
	}
	return true
}

// checks for minimum length string
func (f *Form) CheckFields(fields ...string) {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	// range through the fields
	for _, field := range fields {
		value := strings.TrimSpace(f.Get(field))
		if value != "" {
			// minimum string length
			if len(value) < 3 {
				f.Errors.Add(field, "Gebruik minstens DRIE letters!")
			}
			// check characters only
			if !isAlpha(f.Get(field)) {
				f.Errors.Add(field, "Gebruik ALLÉÉN letters!")
			}
		}
	}
}

func (f *Form) CheckIsEmail(field string) {
	val := strings.TrimSpace(f.Get(field))
	_, err := mail.ParseAddress(val)
	if err != nil {
		f.Errors.Add(field, "Dit lijkt geen email adres!")
	}
}

func (f *Form) EmailExists(field string) {
	f.Errors.Add(field, "Dit emailadres is al in gebruik!")
}
// create a variatic function, we can put as many strings
// as we want is parameters.
func (f *Form) IfFieldsEmpty(fields ...string) bool {
	// we range trough the fields variable
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) != "" {
			for _, field := range fields {
				f.Errors.Add(field, "Mag niet leeg zijn!")
			}
			return false
		}
	}
	return true
}
// checks for all characters string
func (f *Form) CheckAllChars(field string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	// check characters only
	return !isAlpha(f.Get(field))
}
// checks for all characters string
func (f *Form) CheckDigitsOnly(field string) bool {
	isDigit := regexp.MustCompile(`^[1-9]+$`).MatchString
	// check digits only
	return isDigit(f.Get(field))
}
// checks permission box checked
func (f *Form) CheckPermissionBox(field string) bool {
	if f.Get(field) == "on" {
		return true
	} else {
		f.Errors.Add(field, "Toestemming geven is noodzakelijk!")
		return false
	}
}

// checks whether required field is in post and not empty
// "r" is what the browser send back - in the post request
// we can use this for example with check boxes
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		// f.Errors.Add(field, "Beiden leeg, kan niet!")
		return false
	}
	return true
}
