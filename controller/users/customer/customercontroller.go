package customer

import (
	"OKVS2/config"
	login2 "OKVS2/domain/login"
	"OKVS2/domain/users"
	"OKVS2/io/login"
	customerIO "OKVS2/io/users/customer"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

type Customerse users.Customer

func Customer(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/home", CustomerMethod(app))
	r.Get("/table", CustomerTableHandler(app))
	r.Get("/register/{pasword}", RegisterCustomerTableHandler(app))
	return r
}

func RegisterCustomerTableHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pasword := chi.URLParam(r, "pasword")
		logindetails, err := login.GetUserEmail(pasword)
		if err != nil {

		}
		type PageData struct {
			Entities login2.Login
		}
		data := PageData{logindetails}
		files := []string{
			app.Path + "customertable.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func CustomerTableHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var dust []customerIO.Customer
		resp, err := customerIO.GetCustomers()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []customerIO.Customer
		}

		data := PageData{resp}
		files := []string{
			app.Path + "customertable.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

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
