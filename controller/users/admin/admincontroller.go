package admin

import (
	"OKVS2/config"
	items2 "OKVS2/domain/items"
	"OKVS2/io/items"
	users_io "OKVS2/io/users_io/address"
	customer2 "OKVS2/io/users_io/customer"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"html/template"
	"net/http"
)

func Admin(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/home", AdminMethod(app))
	r.Get("/table", AdminTableHandler(app))
	r.Get("/addSpecials", AdminAddSpecialsHandler(app))
	r.Get("/getcustomer/{customerId}", AdminGetCustomerHandler(app))
	///create/special
	r.Post("/create/special", AdminCreateSpecialsHandler(app))
	return r
}

func AdminCreateSpecialsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		itemId := r.PostFormValue("itemId")
		specialType := r.PostFormValue("specialType")
		price := r.PostFormValue("price")
		description := r.PostFormValue("description")
		endDate := r.PostFormValue("endDate")
		file := r.PostFormValue("file")

	}
}

func AdminAddSpecialsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")

		fmt.Println(userEmail)

		//if userEmail == "" {
		//	fmt.Println("error the userEmail is empty",userEmail)
		//	app.ErrorLog.Println("User need to logIn")
		//	http.Redirect(w, r, "/user/login", 301)
		//}
		//admin, err := admin2.GetAdmin(userEmail)
		//if err != nil {
		//	fmt.Println("error the reading admin",admin)
		//	app.ErrorLog.Println("User need to logIn as an Admin")
		//	http.Redirect(w, r, "/user/login", 301)
		//}
		//if admin.Email == "" {
		//
		//	fmt.Println("error the reading admin.Email",admin.Email)
		//	app.ErrorLog.Println("User need to logIn as an Admin")
		//	http.Redirect(w, r, "/user/login", 301)
		//}
		specialType, err := items.GetSpecialTypes()
		if err != nil {
			fmt.Println("error reading customer specialType>>>>", specialType)
			app.ErrorLog.Println(err.Error())
		}
		items, err := items.GetProducts()
		if err != nil {
			fmt.Println("error reading customer items>>>>", items)
			app.ErrorLog.Println(err.Error())
		}
		type PageData struct {
			SpecialType []items2.SpecialType
			Items       []items2.Products
		}
		data := PageData{specialType, items}

		files := []string{
			//app.Path + "itemAdd/addSpecials.html",
			//app.Path + "items/itemProduct.html",C:\Users\ESPOIR\GolandProjects\OKVS2\views\html\itemAdd\addSpacials.html
			app.Path + "itemAdd/addSpacials.html",
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

type customerData struct {
	CustomerName    string
	CustomerSurName string
	Statust         string
	Address         string
	PhoneNumber     string
}

func AdminGetCustomerHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cust customerData
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		fmt.Println(userEmail)

		customerId := chi.URLParam(r, "customerId")

		fmt.Println("customer email: ", customerId)

		custAddress, err := users_io.GetAddress(customerId)
		if err != nil {
			fmt.Println("error reading customer address>>>>", custAddress)
			app.ErrorLog.Println(err.Error())
		}
		customer, err := customer2.GetCustomer(customerId)
		if err != nil {
			fmt.Println("error reading customer>>>>", customer)
			app.ErrorLog.Println(err.Error())
		}
		fmt.Println("customer data: ", customer)
		fmt.Println("custAddress data: ", custAddress)
		if custAddress.Address != "" || customer.Name != "" {
			cust = customerData{customer.Name, customer.SurName, customer.Status, custAddress.Address, custAddress.PhoneNumber}
		}
		fmt.Println("customer>>>>", cust)
		render.JSON(w, r, cust)
	}

}

func AdminTableHandler(app *config.Env) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		type PageData struct {
			name string
		}
		//data := PageData{""}

		files := []string{
			app.Path + "admin/admintable.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(writer, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
func AdminMethod(app *config.Env) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		type PageData struct {
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "put the paths here",
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
