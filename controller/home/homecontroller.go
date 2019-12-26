package home

import (
	"OKVS2/config"
	"OKVS2/domain/items"
	"OKVS2/domain/users"
	"OKVS2/io/login"
	"OKVS2/io/makeUp"
	"OKVS2/io/order"
	"OKVS2/io/users_io/customer"
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
		we first collect all the items that should appear on the home page
		if any thing hapens we send the tamplete home page
		we need to find out the data from the session so that we can che if the user has a card
		*/
		var itemsdetals []items.ItemViewHtml

		homePageElements, err := makeUp.GetAllItems()
		//fmt.Println("User may not have logIn or may not have ordered yet ", homePageElements)
		if err != nil && homePageElements == nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/homeError/homeError", 301)
		}

		//reading the session
		userEmail := app.Session.GetString(r.Context(), "userEmail")

		var message string
		var class string
		var Manager bool

		fmt.Println("User email from the session>>: ", userEmail)
		user, err := customer.GetCustomer(userEmail)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			userLog, err := login.GetUserWithEmail(user.Email)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				if userLog.UserTupe == "admin" {
					Manager = true
					http.Redirect(w, r, "/user/managementwelcom", 301)
					return
				} else {
					Manager = false
				}
			}

		}

		//Checking the card table if there something for this User we will send a message and set a trolley color to danger
		cardDetails, err := order.GetCardWithCustId(userEmail)
		fmt.Println("User card>>: ", cardDetails)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println("User may not have logIn or may not have ordered yet ")
		}
		var itemIdfromcard string
		for _, valeu := range cardDetails {
			itemIdfromcard = valeu.ItemId
		}

		if itemIdfromcard != "" {
			//app.ErrorLog.Println(err.Error())
			message = "You have something in your Card please click on the trolley icon to view your card"
			class = "primary"
		}
		println("homePageElements:  ", homePageElements)

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
		}
		data1 := CardeData{message, class}
		data := PageData{itemsdetals, data1, MyUser{userEmail}, Manager}
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
