package customer

import (
	"OKVS2/config"
	gender2 "OKVS2/domain/gender"
	"OKVS2/domain/items"
	login2 "OKVS2/domain/login"
	"OKVS2/domain/users"
	gender3 "OKVS2/io/gender"
	"OKVS2/io/login"
	"OKVS2/io/makeUp"
	"OKVS2/io/order"
	"OKVS2/io/types"
	address2 "OKVS2/io/users_io/address"
	"OKVS2/io/users_io/customer"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

type CardeData struct {
	Mesage string
	Class  string
}
type MyUser struct {
	User string
}
type Customerse users.Customer

func Customer(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/home", CustomerMethod(app))
	r.Get("/table", CustomerTableHandler(app))
	r.Get("/register/{pasword}", RegisterCustomerHandler(app))
	r.Post("/myregistration", CustomerRegistration(app))
	r.Post("/create/address", CreateAddressHandler(app))
	return r
}

func CreateAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
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
		r.ParseForm()
		genderData := r.PostFormValue("gender")
		surname := r.PostFormValue("surname")
		name := r.PostFormValue("name")
		addressType := r.PostFormValue("addressType")
		cellphone := r.PostFormValue("cellphone")
		address := r.PostFormValue("address")
		age := r.PostFormValue("age")

		if genderData != "" || age != "" {
			readgender, _ := gender3.ResdWithGender(genderData)
			customerGenderObj := gender2.CustomerGender{userEmail, readgender.GenderId, age}
			customerGender, err := customer.CreateCustomerGender(customerGenderObj)
			if err != nil {
				fmt.Println("error in creating customerGender>>: ", customerGender)
				app.ErrorLog.Println(err.Error())
			}

			if addressType != "" || address != "" || cellphone != "" {
				addressobj, _ := types.ReadWithAddressType(addressType)

				addressObj := users.AddressHelper{"00", userEmail, address, addressobj.AddressTypeId, cellphone}
				customerAddress, err := address2.CreateAddress(addressObj)
				if err != nil {
					fmt.Println("error in creating address>>: ")
					app.ErrorLog.Println(err.Error())
				}
				if name != "" || surname != "" {
					customerDetails := users.Customer{userEmail, name, surname, "active"}
					newcustomerDetails, err := customer.UpdateCustomer(customerDetails)
					if err != nil {
						fmt.Println("error in creating address>>: ")
						app.ErrorLog.Println(err.Error())
					}
				}
			}
		}
	}
}

func CustomerRegistration(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		password1 := r.PostFormValue("password1")
		password2 := r.PostFormValue("password2")
		fmt.Println("new password1>>: ", password1+"new password1>>: ", password2+"  userEmail", email)
		if password1 != password2 {
			fmt.Println("new password1>>: ", password1+"new password1>>: ", password2)
			logindetails, err := login.GetUserWithEmail(email)
			customerdetails, _ := customer.GetCustomer(logindetails.Email)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				homeHanler(app)
			}
			type PageData struct {
				Entities login2.Login
				Customer users.Customer
				Class    string
				Message  string
			}
			data := PageData{logindetails, customerdetails, "danger", "please check if your password are the same"}
			files := []string{
				app.Path + "customerUser/passwordUpdate.html",
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
		} else {
			logindetails, err := login.GetUserWithEmail(email)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				//homeHanler(app)
			}
			newLoging := login2.Login{logindetails.Email, password1, logindetails.UserTupe}
			result, _ := login.UpdateLogin(newLoging)
			fmt.Println("user login new details>>: ", result)

			http.Redirect(w, r, "/", 301)
			return
		}
	}
}

func RegisterCustomerHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pasword := chi.URLParam(r, "pasword")
		fmt.Println("generated pasword>>: ", pasword)
		logindetails, err := login.GetUserEmail(pasword)
		fmt.Println("user login details>>: ", logindetails)
		customerdetails, _ := customer.GetCustomer(logindetails.Email)
		fmt.Println("user login customerdetails>>: ", customerdetails)
		if err != nil {
			homeHanler(app)
		}
		type PageData struct {
			Entities login2.Login
			Customer users.Customer
			Class    string
			Message  string
		}
		data := PageData{logindetails, customerdetails, "", ""}
		files := []string{
			app.Path + "customerUser/passwordUpdate.html",
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
			http.Redirect(w, r, "/home/homeError/homeError", 301)
			return
		}

		//reading the session
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		var message string
		var class string

		fmt.Println("User email from the session>>: ", userEmail)
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
		}
		data1 := CardeData{message, class}
		data := PageData{itemsdetals, data1, MyUser{userEmail}}
		files := []string{
			app.Path + "errorPage/error.html",
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
func readImage(byteImage []byte) string {
	mybyte := string(byteImage)
	return mybyte
}
func CustomerTableHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var dust []customerIO.Customer
		resp, err := customer.GetCustomers()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []users.Customer
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
