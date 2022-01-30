package views

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// M is an alias for map[string]interface{}
type M map[string]interface{}

// ErrorResponse represents an error message to be returned
type ErrorResponse struct {
	StatusCode int          `json:"status_code"`
	Detail     string       `json:"detail"`
	Fields     []ErrorField `json:"fields"`
}

// ErrorField represents a field validation error to be returned
type ErrorField struct {
	Location string `json:"location"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
}

// JSON writes the given status code and interface to w as a JSON object
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	enc.Encode(data)
}

// Error writes the given status code and detail string to w as JSON
func Error(w http.ResponseWriter, status int, detail string) {
	res := ErrorResponse{
		StatusCode: status,
		Detail:     detail,
		Fields:     []ErrorField{},
	}
	JSON(w, status, res)
}

// ErrorWithFields writes the given status code and detail string, along with a list
// of `ErrorField`s to w as JSON
func ErrorWithFields(w http.ResponseWriter, status int, detail string, fields []ErrorField) {
	res := ErrorResponse{
		StatusCode: status,
		Detail:     detail,
		Fields:     fields,
	}
	JSON(w, status, res)
}

// DecodeJSON decodes the JSON contents of r into the data interface given
func DecodeJSON(r io.Reader, data interface{}) error {
	dec := json.NewDecoder(r)
	if err := dec.Decode(&data); err != nil {
		return err
	}
	return nil
}

// Validate returns a slice of `ErrorField`s containing any validaton errors on
// the given interface and returns nil if there aren't any
func Validate(data interface{}) []ErrorField {
	// Create new validator and have it use json tags for field names
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Validate the struct and construct a slice of `ErrorField`s
	err := validate.Struct(data)
	if err != nil {
		var errors []ErrorField
		for _, err := range err.(validator.ValidationErrors) {
			var detail string
			if len(err.Param()) > 0 {
				detail = err.ActualTag() + ": " + err.Param()
			} else {
				detail = err.ActualTag()
			}

			errors = append(errors, ErrorField{
				Location: err.Field(),
				Type:     err.Type().String(),
				Detail:   detail,
			})
		}

		return errors
	}

	return nil
}
