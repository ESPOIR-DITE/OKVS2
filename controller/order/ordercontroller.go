package order

import (
	"OKVS2/config"
	"OKVS2/domain/orders"
	"OKVS2/io/order"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func Order(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/addToCard/{resetkeys}", AddToCardHandler(app))
	r.Get("/home", OrderTableHandler(app))
	r.Get("/read/card", ReadCardHandler(app))

	return r
}

func ReadCardHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		cardList, _ := order.GetCardWithCustId(userEmail)

	}
}

type CardeData struct {
	Mesage string
	Class  string
}

func AddToCardHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		itemId := chi.URLParam(r, "resetkeys")
		var message string
		var class string

		fmt.Println("product id to add to the card>>>", itemId, "<<<< User email>>>", userEmail)

		if userEmail == "" {
			app.ErrorLog.Println("User need to logIn")
			http.Redirect(w, r, "/user/login", 301)
			return
		}
		makeCard := orders.Card{"", itemId, userEmail}
		card, err := order.CreateCard(makeCard)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			message = "An error has occured please try again"

		}
		if card.ItemId == "" {
			app.ErrorLog.Println("card.ItemId is empty")
			message = "You have an Item in your Card"
			class = "warning"

		}

		type PageData struct {
			Entity CardeData
		}
		data1 := CardeData{message, class}
		data := PageData{data1}
		fmt.Println(data)
		http.Redirect(w, r, "/", 301)
		return
		/**
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
		}*/
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
