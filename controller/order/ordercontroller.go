package order

import (
	"OKVS2/config"
	"OKVS2/domain/items"
	"OKVS2/domain/orders"
	"OKVS2/domain/users"
	items2 "OKVS2/io/items"
	"OKVS2/io/order"
	customer2 "OKVS2/io/users_io/customer"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"strconv"
)

func Order(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/addToCard/{resetkeys}", AddToCardHandler(app))
	r.Get("/table", OrderTableHandler(app))
	//r.Get("/home", OrderTableHandler(app))
	r.Get("/order/readCard", ReadCardHandler(app))
	r.Get("/order/myorder", MyOrderHandler(app))
	r.Post("/card/item", AddItemToCardHandler(app))
	r.Get("/addToCard/remove/{toremove}", CardRemoveHandler(app))
	r.Get("/track", OrderTrackingHandler(app))

	r.Post("/mytracking", MyTrackingHandler(app))
	return r
}

func MyTrackingHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		fmt.Println("in MyTrackingHandler user id to add to the card>>><<<< User email>>>", userEmail)

		var message string
		if userEmail == "" {
			files := []string{
				app.Path + "loginpage.html",
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
		//need o check if the user exist in the database
		entit, err := customer2.GetCustomer(userEmail)
		if err != nil {
			fmt.Println("in MyTrackingHandler user id to add to the card>>><<<< customer>>>", entit)
			files := []string{

				app.Path + "loginpage.html",
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
		r.ParseForm()
		orderNumber := r.PostFormValue("orderNumber")
		fmt.Println("in MyTrackingHandler user id to add to the card>>><<<< orderNumber>>>", orderNumber)

		entity := orders.OrderHelper{}
		if orderNumber != "" {
			entity, err = order.OrderTracking(orderNumber)
			if err != nil {
				message = "Wrong OrderNumber please try again"
			}
		}
		type PageData struct {
			Entity orders.OrderHelper
			Mesage string
		}
		Data := PageData{entity, message}

		if userEmail != "" {
			files := []string{
				app.Path + "orderPages/order_tracking.html",
			}
			ts, err := template.ParseFiles(files...)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				return
			}
			err = ts.Execute(w, Data)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
	}
}

func OrderTrackingHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		fmt.Println("in OrderTrackingHandler user id to add to the card>>><<<< User email>>>", userEmail)

		if userEmail == "" {
			files := []string{
				app.Path + "loginpage.html",
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

		files := []string{
			app.Path + "tracking.html",
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

func MyOrderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		var message string
		var class string
		var result bool
		fmt.Println("product id to add to the card>>><<<< User email>>>", userEmail)

		if userEmail == "" {
			app.ErrorLog.Println("User need to logIn")
			http.Redirect(w, r, "/user/login", 301)
			return
		}

		//reading all the card of this user
		card, err := order.GetCardWithCustId(userEmail)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			message = "An error has occured please try again"
		}
		fmt.Println("Result after remove ", card)
		for _, cardResult := range card {
			result, _ = order.CreateOrder(cardResult)
		}
		if result != false {

			message = "You have placed an order"
		}
		fmt.Println("Result after placing the order ", result)

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
		//fmt.Println("<<<< User email>>>", userEmail)
		var check []orders.CheckOutHelper
		/**if userEmail == "" {
			app.ErrorLog.Println("User need to logIn")
			http.Redirect(w, r, "/user/login", 307)
			return
		}*/
		cardList, _ := order.GetCardWithCustId(userEmail)
		//fmt.Println("product id to add to the card>>>", cardList, "<<<< User email>>>", userEmail)

		for _, card := range cardList {
			cardCheck, _ := order.GetCheckOut(card)
			check = append(check, orders.CheckOutHelper{readImage(cardCheck.Image), cardCheck.Description, cardCheck.Price, cardCheck.Quantity, cardCheck.Total, cardCheck.ItemId})
		}

		//fmt.Println("product id to add to the card>>>", cardList, "<<<< User email>>>", userEmail)
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

type orderDateils struct {
	Order      order.Order
	Customer   users.Customer
	OderStatus orders.OrderStatus
	OrderLine  []orders.OrderLine
	Items      []myItem
}
type myItem struct {
	Item     items.Products
	Price    float64
	Quantity float64
}

func OrderTableHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myOrderLine []orders.OrderLine
		var itemlist []myItem
		var orderData []orderDateils
		type PageData struct {
			OrderDetails []orderDateils
		}
		orders, err := order.GetOrders()
		if err != nil {
			fmt.Println("error reading orsers in OrderTableHandler")
			app.ErrorLog.Println(err.Error())
		}
		if orders != nil {
			for _, myorder := range orders {
				fmt.Println("my customer>>>", myorder)
				customer, err := customer2.GetCustomer(myorder.CustomerId)
				if err != nil {
					fmt.Println("error reading orsers in OrderTableHandler")
					app.ErrorLog.Println(err.Error())
				}
				myOrderLine, err = order.GetOrderLineWithOrderId(myorder.Id)
				if err != nil {
					fmt.Println("error reading orsers in OrderTableHandler")
					app.ErrorLog.Println(err.Error())
				} else {
					for _, orderL := range myOrderLine {
						product, err := items2.GetProduct(orderL.ItemNumber)
						if err != nil {
							fmt.Println("error reading items in OrderTableHandler")
							app.ErrorLog.Println(err.Error())
						}
						account, err := items2.GetAccounting(product.Id)
						if account.Price != 0 {
							price := account.Price * orderL.Quantity
							itemlist = append(itemlist, myItem{product, price, orderL.Quantity})
						}
					}
				}
				orderStatus, err := order.GetWithOrderId(myorder.Id)

				orderData = append(orderData, orderDateils{myorder, customer, orderStatus, myOrderLine, itemlist})

				fmt.Println("orderData>>>> ", orderData)
			}
		}

		data := PageData{orderData}

		files := []string{
			app.Path + "/admin/ordertable.html",
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
