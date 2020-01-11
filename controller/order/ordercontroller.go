package order

import (
	"OKVS2/config"
	"OKVS2/domain/items"
	"OKVS2/domain/orders"
	"OKVS2/domain/users"
	items2 "OKVS2/io/items"
	"OKVS2/io/order"
	admin2 "OKVS2/io/users_io/admin"
	customer2 "OKVS2/io/users_io/customer"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"strconv"
	"time"
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
	r.Post("/update", UpdateOrderHandler(app))
	return r
}

func UpdateOrderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")

		var notice string
		var class string

		if userEmail == "" {
			notice = "Could not update because there is an error with your identities"
			class = "danger"
			app.ErrorLog.Println("User need to logIn")
			http.Redirect(w, r, "/user/login", 301)
			return
		}
		admin, err := admin2.GetAdmin(userEmail)
		if err != nil {
			notice = "Could not update because there is an error"
			class = "danger"
			app.ErrorLog.Println("User need to logIn as an Admin")
			http.Redirect(w, r, "/user/login", 301)
			return
		}

		if admin.Email == "" {
			notice = "Could not update because there is an error you should login with admin details"
			class = "danger"
			app.ErrorLog.Println("User need to logIn as an Admin")
			http.Redirect(w, r, "/user/login", 301)
			return
		}

		r.ParseForm()

		theOrderId := r.PostFormValue("theOrderId")
		statId := r.PostFormValue("statId")
		orderStatus := orders.OrderStatus{"", theOrderId, time.Now(), userEmail, statId}
		newOrderStatus, err := order.CreateOrderStatus(orderStatus)
		if err != nil {
			notice = "Could not update because there is an error"
			class = "danger"
			fmt.Println("error creating newOrderStatus", newOrderStatus)
			app.ErrorLog.Println(err.Error())
		}
		notice = "You have successfully updated order stat"
		class = "success"

		app.Session.Put(r.Context(), "notice", notice)
		app.Session.Put(r.Context(), "class", class)
		http.Redirect(w, r, "/order/table", 301)
		return
	}
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
	Order    order.Order
	Customer users.Customer
	//OderStatus []orders.OrderStatus
	//OrderLine  []orders.OrderLine
	Items     []myItem
	Status    []orders.Status
	OrderStat []TheOrderStat
}

type myItem struct {
	Item     items.Products
	Price    float64
	Quantity float64
}

func getOrder(orderLine orders.OrderLine) myItem {
	//var itemlist=make( []myItem,0)
	entity := myItem{}
	product, err := items2.GetProduct(orderLine.ItemNumber)
	if err != nil {
		fmt.Println("error reading items in getOrder")
	}
	account, err := items2.GetAccounting(product.Id)
	if err == nil {
		price := account.Price * orderLine.Quantity
		return myItem{product, price, orderLine.Quantity}
	}
	return entity
}

type TheOrderStat struct {
	OrderId    string
	Date       string
	ModifiedBy string
	Stat       string
}

func OrderTableHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		notice := app.Session.GetString(r.Context(), "notice")
		class := app.Session.GetString(r.Context(), "class")
		fmt.Println(userEmail)

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
			//fmt.Println("error the userEmail is empty",userEmail)
			//app.ErrorLog.Println("User need to logIn")
			//http.Redirect(w, r, "/user/login", 301)
		}
		admin, err := admin2.GetAdmin(userEmail)
		if err != nil {
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
			//fmt.Println("error the reading admin",admin)
			//app.ErrorLog.Println("User need to logIn as an Admin")
			//http.Redirect(w, r, "/user/login", 301)
		}

		if admin.Email == "" {
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
			//fmt.Println("error the reading admin.Email",admin.Email)
			//app.ErrorLog.Println("User need to logIn as an Admin")
			//http.Redirect(w, r, "/user/login", 301)
		}

		var theOrderLine []orders.OrderLine
		var theCustomer users.Customer
		var theorderDateils []orderDateils
		var statList []TheOrderStat
		var Order order.Order
		var OderStatus []orders.OrderStatus
		var theLocalItemList []myItem
		statusList, err := order.GetStatues()

		if err != nil {
			fmt.Println("err reading statusList", statusList)
			app.ErrorLog.Println(err.Error())
		}
		ordersList, err := order.GetOrders()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

		//looping through each order to get the following details: 1) order . 2)orderLine
		for _, myOrder := range ordersList {

			//1) getting the Order details
			Order = myOrder

			//2)orderLine details with the orderNumber
			theOrderLine, err = order.GetOrderLineWithOrderId(myOrder.Id)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				return
			}
			theCustomer, err = customer2.GetCustomer(myOrder.CustomerId)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				return
			}

			OderStatus, err = order.GetAllFor(myOrder.Id)
			if err != nil {
				fmt.Println("err reading OderStatus", OderStatus)
				app.ErrorLog.Println(err.Error())
			}
			for _, orderSt := range OderStatus {
				stat, err := order.GetStatus(orderSt.Stat)
				if err != nil {
					fmt.Println("err reading stat", stat)
					app.ErrorLog.Println(err.Error())
				}
				thedate := getDate_YYYYMMDD(orderSt.Date.String())
				fmt.Println("thedate>>>", thedate)
				theStat := TheOrderStat{orderSt.OrderId, thedate, orderSt.ModifiedBy, stat.Stat}
				statList = append(statList, theStat)
			}

			for _, orderLing := range theOrderLine {

				item, err := items2.GetProduct(orderLing.ItemNumber)
				//fmt.Println(item,"<<<orderLing || index>>>>",index)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					return
				}

				account, err := items2.GetAccounting(item.Id)
				if err == nil {
					price := account.Price * orderLing.Quantity
					itmeobj := myItem{item, price, orderLing.Quantity}
					theLocalItemList = append(theLocalItemList, itmeobj)
				}
				//theProduct=append(theProduct,item)
			}

			theorderDateils = append(theorderDateils, orderDateils{Order, theCustomer, theLocalItemList, statusList, statList})
			//fmt.Println("theCustomer >>>",theCustomer.Email,"\ntheProduct",theLocalItemList," \ntheAccount",theAccount," \n\n\n")
			theLocalItemList = nil
			theOrderLine = nil
			statList = nil
		}

		//fmt.Println("theCustomer >>>",theorderDateils)
		//for index,dataorder:=range theorderDateils{
		//	fmt.Println("index: ",index,"    theCustomer >>>",dataorder.Customer,"      order: ",dataorder.Items,"    ",dataorder.OderStatus)
		//}

		type PageData struct {
			OrderDetails []orderDateils
			StatusL      []orders.Status
			Notice       string
			Class        string
		}
		data := PageData{theorderDateils, statusList, notice, class}
		files := []string{
			app.Path + "/admin/ordertable.html",
			app.Path + "template/admin_navbar.html",
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
func getDate_YYYYMMDD(dateString string) string {
	//return strings.Split(dateString, " ")[0]
	layout := "Mon Jan 02 2006 15:04:05 GMT-0700"
	str := dateString
	t, _ := time.Parse(layout, str)
	return t.Format(layout)
}
