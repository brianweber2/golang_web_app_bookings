package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom Form struct, embeds url.Values object
type Form struct {
	URLValues url.Values
	Errors    errors
}

// New initializes a Form struct
func New(data url.Values) *Form {
	return &Form{
		URLValues: data,
		Errors:    errors(map[string][]string{}),
	}
}

// Valid returns true if there are no errors otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}
