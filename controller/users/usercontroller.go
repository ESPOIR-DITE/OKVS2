package users

import (
	"OKVS2/config"
	"OKVS2/domain/users"
	"OKVS2/io/users_io"
	login2 "OKVS2/io/users_io/login"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
	Info  string
}

func User(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/register", userRegisterHandler(app))
	r.Get("/login", userLoginHandler(app))
	r.Get("/relogin", userReloginHandler(app))

	r.Get("/managementwelcom", ManagementHandler(app))
	r.Get("/management", ManagementLoginHandler(app))
	r.Get("/logout", LogoutHandler(app))

	r.Post("/customer/create", CreateCustomerHandler(app))
	r.Post("/customer/log", CustomerLogHandler(app))
	r.Post("/manager/log", ManagerLogHandler(app))

	return r
}

func userReloginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Session.Destroy(r.Context())

		http.Redirect(w, r, "/user/login", 301)
		return
	}
}

func LogoutHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Session.Destroy(r.Context())

		http.Redirect(w, r, "/", 301)
		return
	}
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
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		userLog, err := login2.GetUserWithEmail(userEmail)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/user/management", 301)
			return
		} else if userLog.UserType != "admin" {
			app.Session.Put(r.Context(), "loging", "Wrong Credentials!")
			http.Redirect(w, r, "/user/login", 301)
			return
		}

		files := []string{
			app.Path + "/admin/welcommanagement.html",
			app.Path + "template/admin_navbar.html",
			app.Path + "template/admin_toolbarTemplate.html",
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

		logingDetails := users.LoginHelper{email, password}
		customerDetails := users.Login{logingDetails.Email, logingDetails.Pasword, "customer"}
		resp, err := login2.UserLogin(customerDetails)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			app.Session.Put(r.Context(), "message", "Wrong Credentials!")
			http.Redirect(w, r, "/user/management", 301)
		}
		type PageData struct {
			LoginStat string
		}
		fmt.Println("user type is: ", resp)
		if resp.UserType == "admin" {
			app.Session.Cookie.Name = "UserID"
			app.Session.Put(r.Context(), "userEmail", logingDetails.Email)
			app.Session.Put(r.Context(), "password", logingDetails.Pasword)
			app.InfoLog.Println("Login is successful. Result is ", logingDetails)
			http.Redirect(w, r, "/user/managementwelcom", 301)
		}
		if resp.UserType != "admin" {
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
		fmt.Println("email: ", email+" password: ", password)

		userLoginDetails := users.Login{email, password, ""}
		resp, err := login2.UniversalLogin(userLoginDetails)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			app.Session.Put(r.Context(), "message", "Wrong Credentials!")
			http.Redirect(w, r, "/user/login", 301)
			return
		}
		app.Session.Cookie.Name = "UserID"
		app.Session.Put(r.Context(), "userEmail", resp.Email)
		app.Session.Put(r.Context(), "password", resp.Password)
		app.InfoLog.Println("Login is successful. Result is ", resp)

		if resp.UserType == "admin" {
			http.Redirect(w, r, "/user/managementwelcom", 301)
		} else {
			http.Redirect(w, r, "/", 301)
		}
	}
}

func CreateCustomerHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" we are creating ")
		r.ParseForm()
		name := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		email := r.PostFormValue("email")
		user := users.Customer{email, name, lastName, "active"}

		result, err := users_io.CreateCustomer(user)
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

func userRegisterHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "customer/customerform.html",
			app.Path + "template/navigator.html",
			app.Path + "template/footer.html",
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
