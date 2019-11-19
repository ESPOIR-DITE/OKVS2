package home

import (
	"OKVS2/config"
	"OKVS2/domain/items"
	"OKVS2/domain/users"
	"OKVS2/io/makeUp"
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
	r.Get("/", homeHanler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func indexErrorHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		println("excuting indexErrorHanler:  ")
		files := []string{
			app.Path + "index2.html",
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

type ImageItems struct {
	Pic string
}

// this method help for converting []byte to strings
func readImage(byteImage []byte) string {
	mybyte := string(byteImage)
	return mybyte
}
func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var itemsdetals []items.ItemViewHtml
		homePageElements, err := makeUp.GetAllItems()
		println("homePageElements:  ", homePageElements)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/home/homeError/homeError", 301)
			return
		}
		if homePageElements != nil {
			for _, itemImageId := range homePageElements {
				itemsdetals = append(itemsdetals, items.ItemViewHtml{itemImageId.ItemNumber, itemImageId.ProductName, itemImageId.Price, itemImageId.Description, readImage(itemImageId.Image)})
			}
		}
		type PageData struct {
			Entities []items.ItemViewHtml
		}
		data := PageData{itemsdetals}
		files := []string{
			app.Path + "index.html",
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
