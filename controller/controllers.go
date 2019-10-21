package controller

import (
	"OKVS2/config"
	controllers "OKVS2/controller/home"
	"OKVS2/controller/item"
	"OKVS2/controller/users"
	"OKVS2/controller/users/admin"
	"OKVS2/controller/users/customer"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)

	mux.Handle("/", controllers.Home(env))
	mux.Mount("/category", item.Home(env))
	mux.Mount("/customer", customer.Customer(env))
	mux.Mount("/user", users.User(env))
	mux.Mount("/manager", admin.Admin(env))

	fileServer := http.FileServer(http.Dir("./views/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}
