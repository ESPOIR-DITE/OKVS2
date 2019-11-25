package order

import (
	"OKVS2/config"
	"OKVS2/domain/orders"
	"OKVS2/io/order"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"strconv"
)

func Order(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/addToCard/{resetkeys}", AddToCardHandler(app))
	r.Get("/home", OrderTableHandler(app))
	r.Get("/order/readCard", ReadCardHandler(app))
	r.Post("/card/item", AddItemToCardHandler(app))
	r.Get("/addToCard/remove/{toremove}", CardRemoveHandler(app))
	return r
}

func CardRemoveHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		itemId := chi.URLParam(r, "toremove")
		var message string
		var class string

		fmt.Println("product id to add to the card>>>", itemId, "<<<< User email>>>", userEmail)

		if userEmail == "" {
			app.ErrorLog.Println("User need to logIn")
			http.Redirect(w, r, "/user/login", 301)
			return
		}

		makeCard := orders.Card{"", itemId, userEmail, 00}
		card, err := order.RemoveCard(makeCard)

		fmt.Println("Result after remove ", card)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			message = "An error has occured please try again"

		}

		type PageData struct {
			Entity CardeData
		}
		data1 := CardeData{message, class}
		data := PageData{data1}
		fmt.Println(data)
		http.Redirect(w, r, "/order/order/readCard", 301)
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

func AddItemToCardHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")

		var message string
		var class string

		r.ParseForm()
		quantity, _ := strconv.Atoi(r.PostFormValue("qty"))
		itemId := r.PostFormValue("itemId")
		fmt.Println("checking the card>>>", quantity, "<<<< itemId>>>", itemId, "<<<< User email>>>", userEmail)
		if userEmail == "" {
			app.ErrorLog.Println("User need to logIn")
			http.Redirect(w, r, "/user/login", 301)
			return
		}
		makeCard := orders.Card{"", itemId, userEmail, quantity}
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
	}
}
func readImage(byteImage []byte) string {
	mybyte := string(byteImage)
	return mybyte
}

func ReadCardHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		fmt.Println("<<<< User email>>>", userEmail)
		var check []orders.CheckOutHelper
		/**if userEmail == "" {
			app.ErrorLog.Println("User need to logIn")
			http.Redirect(w, r, "/user/login", 307)
			return
		}*/
		cardList, _ := order.GetCardWithCustId(userEmail)
		fmt.Println("product id to add to the card>>>", cardList, "<<<< User email>>>", userEmail)

		for _, card := range cardList {
			cardCheck, _ := order.GetCheckOut(card)
			check = append(check, orders.CheckOutHelper{readImage(cardCheck.Image), cardCheck.Description, cardCheck.Price, cardCheck.Quantity, cardCheck.Total, cardCheck.ItemId})
		}

		fmt.Println("product id to add to the card>>>", cardList, "<<<< User email>>>", userEmail)
		type PageData struct {
			Entity []orders.CheckOutHelper
		}
		//fmt.Println("check >>>", check)
		data := PageData{check}

		files := []string{
			app.Path + "items/item_cart.html",
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

		makeCard := orders.Card{"", itemId, userEmail, 00}
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
