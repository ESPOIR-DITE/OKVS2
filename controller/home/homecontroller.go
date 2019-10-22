package home

import (
	"OKVS2/config"
	"OKVS2/domain/users"
	"OKVS2/io/users/customer"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

type Customer users.Customer
type PageData struct {
	User interface{}
}

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	r.Get("/home", indexHanler(app))
	r.Get("/", homeHanler(app))
	return r
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "index.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func indexHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		fmt.Println(email)

		if email == "" || len(email) <= 0 {
			fmt.Println("email is empty ....")
			files := []string{
				app.Path + "index.html",
			}
			ts, err := template.ParseFiles(files...)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				return
			}
			err = ts.Execute(w, nil)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		userName, err := customer.GetCustomer(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/", 301)
			return
		}

		date := PageData{userName}

		files := []string{
			app.Path + "index.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, date)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}