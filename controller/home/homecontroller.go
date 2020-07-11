package home

import (
	"OKVS2/config"
	helperUser "OKVS2/controller/users"
	"OKVS2/domain/items"
	"OKVS2/domain/users"
	"OKVS2/io/joins"
	"OKVS2/io/users_io"
	"OKVS2/io/users_io/admin"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

type Customer users.Customer
type PageData struct {
	User interface{}
}
type CardeData struct {
	Mesage string
	Class  string
}

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	r.Get("/home", indexHanler(app))
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
			app.Path + "template/footer.html",
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
type User struct {
	TheUser string
}

// this method help for converting []byte to strings
func readImage(byteImage []byte) string {
	mybyte := string(byteImage)
	return mybyte
}

type MyUser struct {
	User string
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/**
		we first collect all the item_io that should appear on the home page
		if any thing hapens we send the tamplete home page
		we need to find out the data from the session so that we can che if the user has a card
		*/

		var itemsdetals []items.ItemViewHtml

		homePageElements, err := joins.GetAllItems()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		//fmt.Println("User may not have logIn or may not have ordered yet ", homePageElements)
		if homePageElements == nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/homeError/homeError", 301)
		}

		//reading the session
		userEmail := app.Session.GetString(r.Context(), "userEmail")

		var message string
		var class string
		var Manager = false
		var user users.Customer
		//var cardDetails []orders.Card

		//([]orders.Card,string,string,bool,users.Customer)
		message, class, Manager, user = helperUser.GetUserDetails(userEmail)
		if Manager == true {
			_, err := admin.GetAdmin(userEmail)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				Manager = true
				http.Redirect(w, r, "/user/managementwelcom", 301)
				return
			}
		}

		if homePageElements != nil {
			for _, itemImageId := range homePageElements {
				itemsdetals = append(itemsdetals, items.ItemViewHtml{itemImageId.ItemNumber, itemImageId.ProductName, itemImageId.Price, itemImageId.Description, readImage(itemImageId.Image)})
			}
		}
		type PageData struct {
			Entities []items.ItemViewHtml
			Entity   CardeData
			MyUser
			Manager bool
			User    users.Customer
		}
		data1 := CardeData{message, class}
		data := PageData{itemsdetals, data1, MyUser{userEmail}, Manager, user}

		files := []string{
			app.Path + "index.html",
			app.Path + "template/footer.html",
			app.Path + "customer-template/toolbarTemplate.html",
			app.Path + "customer-template/navbar.html",
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
				app.Path + "template/navigator.html",
				app.Path + "template/footer.html",
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
		userName, err := users_io.GetCustomer(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/", 301)
			return
		}

		date := PageData{userName}

		files := []string{
			app.Path + "index.html",
			app.Path + "template/navigator.html",
			app.Path + "template/footer.html",
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
