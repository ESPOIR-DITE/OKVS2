package users

import (
	"OKVS2/config"
	"OKVS2/domain/login"
	"OKVS2/domain/users"
	login2 "OKVS2/io/login"
	"OKVS2/io/users/customer"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Title string
	Info  string
}

func User(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/register", userRegisterHandler(app))
	r.Get("/login", userLoginHandler(app))
	r.Get("/managementwelcom", ManagementHandler(app))
	r.Get("/management", ManagementLoginHandler(app))

	r.Post("/customer/create", CreateCustomerHandler(app))
	r.Post("/customer/log", CustomerLogHandler(app))
	r.Post("/manager/log", ManagerLogHandler(app))

	return r
}

func ManagementLoginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := app.Session.GetString(r.Context(), "message")

		var data PageData

		if message != "" {
			data = PageData{"Login Error!", message}
		} else {
			data = PageData{}
		}
		files := []string{
			app.Path + "managementLogin.html",
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

func ManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "welcommanagement.html",
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

func ManagerLogHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var stat string
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		fmt.Println("email: ", email+"password: ", password)

		logingDetails := login.LoginHelper{email, password}
		resp, err := login2.UserLogin(logingDetails)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			app.Session.Put(r.Context(), "message", "Wrong Credentials!")
			http.Redirect(w, r, "/user/management", 301)
		}
		type PageData struct {
			LoginStat string
		}
		fmt.Println("user type is: ", resp)
		if resp.UserTupe == "admin" {
			app.Session.Cookie.Name = "UserID"
			app.Session.Put(r.Context(), "userEmail", logingDetails.Email)
			app.Session.Put(r.Context(), "password", logingDetails.Pasword)
			app.InfoLog.Println("Login is successful. Result is ", logingDetails)
			http.Redirect(w, r, "/user/managementwelcom", 301)
		}
		if resp.UserTupe != "admin" {
			stat = "Please Login here "
		}
		fmt.Println("user type is: ", stat)
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
}

func CustomerLogHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		fmt.Println("email: ", email+"password: ", password)

		//var data PageDate

		logingDetails := login.LoginHelper{email, password}
		resp, err := login2.UserLogin(logingDetails)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			app.Session.Put(r.Context(), "message", "Wrong Credentials!")
			http.Redirect(w, r, "/user/login", 301)
			return
		}
		app.Session.Cookie.Name = "UserID"
		app.Session.Put(r.Context(), "userEmail", email)
		app.Session.Put(r.Context(), "password", password)
		app.InfoLog.Println("Login is successful. Result is ", logingDetails)

		userName, erro := customer.GetCustomer(strings.TrimSpace(resp.Email))

		fmt.Println("the user is ", userName)

		//data :=PageDate{userName.Email,userName.Name}
		if erro != nil {
			fmt.Println("Login fail The Response1>>>", strings.TrimSpace(resp.Email), "<<<<")
			fmt.Println("Login fail The Response2 ", userName)
			http.Redirect(w, r, "/user/login", 301)
		}
		http.Redirect(w, r, "/", 301)
	}
}

func CreateCustomerHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" we are creating ")
		r.ParseForm()
		name := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		email := r.PostFormValue("EMAIL")
		user := users.Customer{email, name, lastName, "active"}

		result, err := customer.CreateCustomer(user)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		app.InfoLog.Println("create response is :", result)
		http.Redirect(w, r, "/", 301)
	}
}

func userLoginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := app.Session.GetString(r.Context(), "message")

		var data PageData

		if message != "" {
			data = PageData{"Login Error!", message}
		} else {
			data = PageData{}
		}
		files := []string{
			app.Path + "loginpage.html",
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

func userRegisterHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "customerform.html",
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
