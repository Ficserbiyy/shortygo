package handler

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v5"
)

var (
	templates = template.Must(template.ParseFiles(
		"index.html",
		"templates/shorten.html",
		"templates/new.html",
		"templates/link.html",
		"templates/edit.html",
	))
)

func renderTemplate(w http.ResponseWriter, tmpl string, p any) error {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
