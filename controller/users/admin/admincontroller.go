package admin

import (
	"OKVS2/config"
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
	r.Get("/getcustomer/{customerId}", AdminGetCustomerHandler(app))
	//r.Get("/table", AdminTableHandler(app))

	return r
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
