package http_server

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func (c *BaseController) WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (c *BaseController) WriteJSONError(w http.ResponseWriter, status int, message string) {
	c.WriteJSON(w, status, map[string]string{"error": message})
}

func (c *BaseController) WriteValidationErrors(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	errors := make(map[string]string)
	for _, e := range validationErrs {
		errors[e.Field()] = e.Tag()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	_ = json.NewEncoder(w).Encode(map[string]any{"errors": errors})
}

// DecodeBody decodes the JSON body into dst (which must be a pointer to a struct),
// validates it with the validator, and indicates success or failure.
// In case of error, writes the appropriate HTTP response and returns false.
func (c *BaseController) DecodeBody(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		c.WriteJSONError(w, http.StatusBadRequest, "invalid JSON body")
		return false
	}

	if err := c.validate.Struct(dst); err != nil {
		c.WriteValidationErrors(w, err)
		return false
	}

	return true
}

// DecodeParams reads URL params (using chi) and populates the struct pointed to by dst.
// It uses the `param` tag, then `json`, and finally the field name in lowerCamelCase.
// After populating, it validates the struct. On error, writes a response and returns false.
func (c *BaseController) DecodeParams(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := bindFromParams(r, dst); err != nil {
		c.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return false
	}

	if err := c.validate.Struct(dst); err != nil {
		c.WriteValidationErrors(w, err)
		return false
	}

	return true
}

// DecodeQuery reads query string parameters and populates the struct pointed to by dst.
// It uses the `query` tag, then `json`, and finally the field name in lowerCamelCase.
// After populating, it validates the struct. On error, writes a response and returns false.
func (c *BaseController) DecodeQuery(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := bindFromQuery(r, dst); err != nil {
		c.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return false
	}

	if err := c.validate.Struct(dst); err != nil {
		c.WriteValidationErrors(w, err)
		return false
	}

	return true
}

var ErrInvalidDestination = errors.New("destination must be a non-nil pointer to a struct")

// bindFromQuery populates the struct pointed to by dst using the query string.
func bindFromParams(r *http.Request, dst any) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return ErrInvalidDestination
	}
	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return ErrInvalidDestination
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !v.Field(i).CanSet() {
			continue
		}

		name := field.Tag.Get("param")
		if name == "" {
			name = field.Tag.Get("json")
		}
		if name == "" {
			name = toLowerCamel(field.Name)
		}

		raw := chi.URLParam(r, name)
		if raw == "" {
			continue
		}

		if err := setFieldValue(v.Field(i), raw); err != nil {
			return err
		}
	}

	return nil
}

// bindFromQuery populates the struct pointed to by dst using the query string.
func bindFromQuery(r *http.Request, dst any) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return ErrInvalidDestination
	}
	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return ErrInvalidDestination
	}

	t := v.Type()
	q := r.URL.Query()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !v.Field(i).CanSet() {
			continue
		}

		name := field.Tag.Get("query")
		if name == "" {
			name = field.Tag.Get("json")
		}
		if name == "" {
			name = toLowerCamel(field.Name)
		}

		raw := q.Get(name)
		if raw == "" {
			continue
		}

		if err := setFieldValue(v.Field(i), raw); err != nil {
			return err
		}
	}

	return nil
}

// setFieldValue converts the raw string to the field's type and assigns it.
// Supports basic types: string, bool, integers.
func setFieldValue(fv reflect.Value, raw string) error {
	ft := fv.Type()

	// Suporte a ponteiros para tipos básicos.
	if ft.Kind() == reflect.Ptr {
		elem := reflect.New(ft.Elem())
		if err := setFieldValue(elem.Elem(), raw); err != nil {
			return err
		}
		fv.Set(elem)
		return nil
	}

	switch ft.Kind() {
	case reflect.String:
		fv.SetString(raw)
	case reflect.Bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return err
		}
		fv.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return err
		}
		fv.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return err
		}
		fv.SetUint(n)
	default:
		// Unsupported type: ignore
	}

	return nil
}

func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = []rune(strings.ToLower(string(runes[0])))[0]
	return string(runes)
}
