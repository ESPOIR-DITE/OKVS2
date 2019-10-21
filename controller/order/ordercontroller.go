package order

import (
	"OKVS2/config"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func Order(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/home", OrderTableHandler(app))
	return r
}

func OrderTableHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type PageData struct {
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "ordertable.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
