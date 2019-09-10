package controller

import (
	"OKVS2/controller/admin"
	"OKVS2/controller/customer"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"obas/src/config"
	controllers "obas/src/controllers/home"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)

	mux.Handle("/", controllers.Home(env))
	mux.Mount("/admin", admin.Admin(env))
	mux.Mount("/customer", customer.Customer(env))

	fileServer := http.FileServer(http.Dir("./src/views/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}
