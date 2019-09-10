package customer

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
)

func Customer(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/home", CustomerMethod(app))
	return r
}
func CustomerMethod(app *config.Env) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		type PageData struct {
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "put the paths here", //ex: app.Path + "/address/address.page.html", of all the html pages that are need for a single interface
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(writer, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
