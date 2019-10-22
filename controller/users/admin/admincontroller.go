package admin

import (
	"OKVS2/config"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func Admin(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/home", AdminMethod(app))
	r.Get("/table", AdminTableHandler(app))
	//r.Get("/table", AdminTableHandler(app))

	return r
}

func AdminTableHandler(app *config.Env) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		type PageData struct {
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "admintable.html",
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
func AdminMethod(app *config.Env) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		type PageData struct {
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "put the paths here",
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
