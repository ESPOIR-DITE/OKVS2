package customer

import (
	"OKVS2/config"
	helperUser "OKVS2/controller/users"
	gender2 "OKVS2/domain/gender"
	"OKVS2/domain/items"
	"OKVS2/domain/users"
	"OKVS2/io/joins"
	"OKVS2/io/order_io/card"
	"OKVS2/io/users_io"
	address2 "OKVS2/io/users_io/address"
	"OKVS2/io/users_io/gender"
	"OKVS2/io/users_io/login"
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
	r.Get("/profile", CustomerProfileHandler(app))
	r.Get("/contact", ContactProfileHandler(app))
	r.Get("/profileEdite", CustomerEditeProfileHandler(app))
	r.Post("/myregistration", CustomerRegistration(app))
	r.Post("/create/address", CreateAddressHandler(app))
	r.Post("/profile/update", UpdateProfileHandler(app))

	return r
}

func ContactProfileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")

		_, err := users_io.GetCustomer(userEmail)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/", 301)
			return
		}

		message, class, Manager, user := helperUser.GetUserDetails(userEmail)

		type PageData struct {
			Entity CardeData
			MyUser
			Manager bool
			User    users.Customer
		}
		data := PageData{CardeData{message, class}, MyUser{userEmail}, Manager, user}

		files := []string{
			app.Path + "customerUser/contactUser.html",
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

func UpdateProfileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		r.ParseForm()
		name := r.PostFormValue("name")
		surName := r.PostFormValue("surname")
		dateOfBirth := r.PostFormValue("deteOfBirth")
		addressTypeId := r.PostFormValue("addressType")
		fmt.Println("addressTypeId ", addressTypeId)
		address := r.PostFormValue("address")
		genderId := r.PostFormValue("gender")
		phoneNumber := r.PostFormValue("cellphone")

		myCustomer := users.Customer{userEmail, name, surName, "active"}
		_, err := users_io.UpdateCustomer(myCustomer)
		if err != nil {
			fmt.Println("myCustomer ", myCustomer)
			app.ErrorLog.Println(err.Error())
		}
		/**getting the id of the gender that**/
		customerGender := gender2.CustomerGender{userEmail, genderId, dateOfBirth}
		_, errr := users_io.UpdateCustomerGender(customerGender)
		fmt.Println("myCustomer Gender ", customerGender)
		if errr != nil {
			fmt.Println("myCustomer Gender ", customerGender)
			app.ErrorLog.Println(errr.Error())
		}
		/**getting the customer address**/
		customerAddress := users.Address{"", userEmail, address, addressTypeId, phoneNumber}
		_, errrr := address2.UpdateAddress(customerAddress)
		if errrr != nil {
			fmt.Println("myCustomer ", customerAddress)
			app.ErrorLog.Println(errrr.Error())
		}
		fmt.Println("customerAddress ", customerAddress)
		fmt.Println("customerGender ", customerGender)
		fmt.Println("myCustomer ", myCustomer)

		http.Redirect(w, r, "/customer/profile", 301)

	}
}

func CustomerEditeProfileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		message, class, Manager, user := helperUser.GetUserDetails(userEmail)

		files := []string{
			app.Path + "/customerUser/profileEdite.html",
			app.Path + "customer-template/toolbarTemplate.html",
			app.Path + "customer-template/navbar.html",
		}
		mycustomer, err := users_io.GetCustomer(userEmail)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/", 301)
			return
		}

		customerGender, err := users_io.GetCustomerGender(userEmail)
		if err != nil {
			fmt.Println("in customerGender")
			app.ErrorLog.Println(err.Error())
		}
		customerAddress, err := address2.GetAddress(userEmail)
		fmt.Println("in customerAddress", customerAddress)
		if err != nil {
			fmt.Println("in customerAddress")
			app.ErrorLog.Println(err.Error())
		}

		genderType, err := gender.GetGenders()
		if err != nil {
			fmt.Println("in genderType")
			app.ErrorLog.Println(err.Error())
		}
		addressTypes, err := address2.GetAddressTypes()

		if err != nil {
			fmt.Print("in addressTypes")
			app.ErrorLog.Println(err.Error())
		}

		type PageData struct {
			AddressTypes    []address2.AddressType
			GenderType      []gender2.Gender
			Customer        users.Customer
			CustomerAddress users.Address
			Gender          gender2.CustomerGender

			Entity CardeData
			MyUser
			Manager bool
			User    users.Customer
		}
		data := PageData{addressTypes, genderType, mycustomer, customerAddress, customerGender, CardeData{message, class}, MyUser{userEmail}, Manager, user}
		//fmt.Print("in data>>>", data)
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

func CustomerProfileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		customerDetail, err := users_io.GetCustomer(userEmail)
		message, class, Manager, user := helperUser.GetUserDetails(userEmail)
		fmt.Println(" in reading customerDetaild>>: customerDetail")
		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println("error in reading customerDetaild>>: ")
		}
		var reportor = "update successeful"
		var Class = "success"
		var checker = true

		customerGender, err := users_io.GetCustomerGender(userEmail)
		fmt.Println("Gender Id>>: ", customerGender)
		if err != nil {
			fmt.Println("error in reading customerGender>>: ")
			app.ErrorLog.Println(err.Error())
		}
		gender, err := gender.GetGender(customerGender.GenderId)
		if err != nil {
			fmt.Println("error in reading Gender>>: ")
			app.ErrorLog.Println(err.Error())
		}
		customerAddress, err := address2.GetAddress(userEmail)
		fmt.Println("error in reading customerAddress>>: ", customerAddress)
		if err != nil {
			fmt.Println("error in reading customerAddress>>: ")
			app.ErrorLog.Println(err.Error())
		}

		type PageDate struct {
			Reportor        string
			Class           string
			Checker         bool
			Customer        users.Customer
			CustomerGender  gender2.Gender
			CustomerAddress users.Address
			Entity          CardeData
			MyUser
			Manager bool
			User    users.Customer
		}
		date := PageDate{reportor,
			Class,
			checker,
			customerDetail,
			gender,
			customerAddress,
			CardeData{message, class},
			MyUser{userEmail},
			Manager,
			user,
		}
		if userEmail == "" {
			files := []string{
				app.Path + "loginpage.html",
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
		files := []string{
			app.Path + "/customerUser/profile.html",
			app.Path + "customer-template/toolbarTemplate.html",
			app.Path + "customer-template/navbar.html",
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

func CreateAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		_, err := users_io.GetCustomer(userEmail)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/", 301)
			return
		}
		var reportor = "update successeful"
		var Class = "success"
		var checker = false
		message, class, Manager, user := helperUser.GetUserDetails(userEmail)

		r.ParseForm()
		genderData := r.PostFormValue("gender")
		surname := r.PostFormValue("surname")
		name := r.PostFormValue("name")
		addressType := r.PostFormValue("addressType")
		cellphone := r.PostFormValue("cellphone")
		address := r.PostFormValue("address")
		age := r.PostFormValue("age")

		if genderData != "" || age != "" {
			readgender, _ := gender.GetGender(genderData)
			customerGenderObj := gender2.CustomerGender{userEmail, readgender.GenderId, age}
			customerGender, err := users_io.CreateCustomerGender(customerGenderObj)
			if err != nil {
				fmt.Println("error in creating customerGender>>: ", customerGender)
				app.ErrorLog.Println(err.Error())
				reportor = "An error has occured"
				Class = "danger"
				checker = false
			}
			if addressType != "" || address != "" || cellphone != "" {
				addressobj, _ := address2.ReadWithAddressType(addressType)
				addressObj := users.UserAddress{"00", userEmail, address, addressobj.AddressTypeId, cellphone}
				_, err := address2.CreateAddress(addressObj)
				if err != nil {
					fmt.Println("error in creating address>>: ")
					app.ErrorLog.Println(err.Error())
					reportor = "An error has occured"
					Class = "danger"
					checker = false
				}
				if name != "" || surname != "" {
					customerDetails := users.Customer{userEmail, name, surname, "active"}
					_, err := users_io.UpdateCustomer(customerDetails)
					if err != nil {
						fmt.Println("error in creating customerDetails>>: ")
						app.ErrorLog.Println(err.Error())
						reportor = "An error has occured"
						Class = "danger"
						checker = false
					}
				}
			}
		}
		customerDetail, err := users_io.GetCustomer(userEmail)
		if err != nil {
			fmt.Println("error in reading customerDetaild>>: ")
			app.ErrorLog.Println(err.Error())
		}
		customerGender, err := users_io.GetCustomerGender(userEmail)
		if err != nil {
			fmt.Println("error in reading customerGender>>: ")
			app.ErrorLog.Println(err.Error())
		}
		gender, err := gender.GetGender(customerGender.GenderId)
		if err != nil {
			fmt.Println("error in reading Gender>>: ")
			app.ErrorLog.Println(err.Error())
		}
		customerAddress, err := address2.GetAddress(userEmail)
		if err != nil {
			fmt.Println("error in reading customerAddress>>: ")
			app.ErrorLog.Println(err.Error())
		}
		type PageDate struct {
			Reportor        string
			Class           string
			Checker         bool
			Customer        users.Customer
			CustomerGender  gender2.Gender
			CustomerAddress users.Address
			Entity          CardeData
			MyUser
			Manager bool
			User    users.Customer
		}
		date := PageDate{reportor,
			Class,
			checker,
			customerDetail,
			gender,
			customerAddress,
			CardeData{message, class},
			MyUser{userEmail},
			Manager,
			user,
		}
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
			app.Path + "/customerUser/profile.html",
			app.Path + "customer-template/toolbarTemplate.html",
			app.Path + "customer-template/navbar.html",
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

func CustomerRegistration(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		previousPassword := r.PostFormValue("previousPassword")
		password1 := r.PostFormValue("password1")
		password2 := r.PostFormValue("password2")
		fmt.Println("new password1>>: ", password1+"new password1>>: ", password2+"  userEmail", email)
		if password1 != "" && password2 != "" {
			fmt.Println("they are not empty......")
			if password1 == password2 {
				fmt.Println("they are equal")
				fmt.Println("new password1>>: ", password1+"new password1>>: ", password2)
				logindetails, err := login.GetUserWithEmail(email)
				if err != nil {
					//Todo we need to report with an error message
					fmt.Println("could not find the login details")
					http.Redirect(w, r, "/customer/register/"+previousPassword, 301)
					return
				}
				_, errx := users_io.GetCustomer(logindetails.Email)
				if errx != nil {
					//Todo we need to report with an error message
					app.ErrorLog.Println(errx.Error())
					fmt.Println("could not find the customer details")
					http.Redirect(w, r, "/customer/register/"+previousPassword, 301)
					return
					//homeHanler(app)
				}
				newLoging := users.Login{email, password2, "customer"}
				_, errr := login.UpdateLogin(newLoging)
				if errr != nil {
					//Todo we need to report with an error message
					app.ErrorLog.Println(errr.Error())
					fmt.Println("could not update Loging details")
					http.Redirect(w, r, "/customer/register/"+previousPassword, 301)
					return
					//homeHanler(app)
				}
				http.Redirect(w, r, "/user/login", 301)
				return
				//type PageData struct {
				//	Entities login2.Login
				//	Customer users.Customer
				//	Class    string
				//	Message  string
				//}
				//data := PageData{logindetails, customerdetails, "danger", "please check if your password are the same"}
				//files := []string{
				//	app.Path + "customerUser/passwordUpdate.html",
				//}
				//ts, err := template.ParseFiles(files...)
				//if err != nil {
				//	app.ErrorLog.Println(err.Error())
				//	return
				//}
				//err = ts.Execute(w, data)
				//if err != nil {
				//	app.ErrorLog.Println(err.Error())
				//}
			} else {
				fmt.Println("confirm with empty password")
				http.Redirect(w, r, "/customer/register/"+previousPassword, 301)
				return
			}
		} else {
			fmt.Println("confirm with empty password")
			http.Redirect(w, r, "/customer/register/"+previousPassword, 301)
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
		customerdetails, _ := users_io.GetCustomer(logindetails.Email)
		fmt.Println("user login customerdetails>>: ", customerdetails)
		if err != nil {
			homeHanler(app)
		}
		type PageData struct {
			Entities users.Login
			Customer users.Customer
			Class    string
			Message  string
			Password string
		}
		data := PageData{logindetails, customerdetails, "", "", pasword}
		files := []string{
			app.Path + "customerUser/passwordUpdate.html",
			app.Path + "template/navigator.html",
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

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		/**
		we first collect all the item_io that should appear on the home page
		if any thing hapens we send the tamplete home page
		we need to find out the data from the session so that we can che if the user has a card
		*/
		var itemsdetals []items.ItemViewHtml

		homePageElements, err := joins.GetAllItems()
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
		cardDetails, err := card.GetCardWithCustId(userEmail)
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
		resp, err := users_io.GetCustomers()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []users.Customer
		}

		data := PageData{resp}
		files := []string{
			app.Path + "/admin/customertable.html",
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
