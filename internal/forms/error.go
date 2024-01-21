package forms

// a package that rules the way serverside validation
// is handled

// it's a slice of string because there might be more
// then one error on a given field - input tag "name"
type errors map[string][]string

// field refers to the name attribute in the input tag
// Add adds an error message for a given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// check whether the field has an error string
// and - if so - returns the first error message
func (e errors) Get(field string) string {
	// check whether e[field] holds a value
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	// return the first index on the error slice
	return es[0]
}