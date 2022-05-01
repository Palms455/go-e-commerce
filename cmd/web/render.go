package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Errr string
	IsAutenticated string
	API string
	CSSVersion string
	StripePublishKey string
	StripeSecretKey string
}

var functions = template.FuncMap{
	"formatCurrency": formatCurrency,
}

func formatCurrency(val int) string {
	format_value := float32(val/100)
	return fmt.Sprintf("$%.2f", format_value)
}
//go:embed templates
var templateFS embed.FS
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	td.API = app.config.api
	td.StripeSecretKey = app.config.stripe.secret
	td.StripePublishKey = app.config.stripe.key
	return td

}


func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var tmpl *template.Template
	var err error

	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateMap := app.templateCache[templateToRender]

	if app.config.env == "prod" && templateMap {
		tmpl = app.templateCache[templateToRender]
	} else {
		tmpl, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}
	if td == nil {
		td = &templateData{}
	}
	td = app.addDefaultData(td, r)

	err = tmpl.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	return nil
}


func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var tmpl *template.Template
	var err error
	if len(partials) > 0 {
		for i, name := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", name)
		}
	}
	if len(partials) >0 {
		tmpl, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		tmpl, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)
	}
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = tmpl

	return tmpl, err
}