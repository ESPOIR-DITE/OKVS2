package order

import (
	"OKVS2/config"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func Order(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/home", OrderTableHandler(app))
	r.Get("/addToCard/{resetkeys}", AddToCardHandler(app))
	return r
}

func AddToCardHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		resetKey := chi.URLParam(r, "resetkeys")
		//productTypeId := r.PostFormValue("productpic")
		fmt.Println("product id to add to the card>>>", resetKey, "<<<< User email>>>", userEmail)

		//fmt.Println("product Details to search>>>", productDetails)

		http.Redirect(w, r, "/", 301)
		return
	}
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
