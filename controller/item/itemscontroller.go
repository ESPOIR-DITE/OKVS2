package item

import (
	"OKVS2/config"
	"OKVS2/domain/items"
	itemsIO "OKVS2/io/items"
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Soulier items.Shoes
type Perique items.Hair
type Items items.Items
type ItemSold items.ItemSold
type Cloths items.Cloths
type Beate items.BeautyMakeup

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	r.Get("/", indexHanler(app))
	r.Get("/soulier/table", SoulierItemHanler(app))
	r.Get("/chemise/table", ChemiseItemHanler(app))
	r.Get("/pantalon/table", PantalonHanler(app))
	r.Get("/beate/table", BeauteItemHanler(app))
	r.Get("/perique/table", PeriqueItemHanler(app))

	r.Get("/soulier/add", SoulierAddHandler(app))
	r.Get("/chemise/add", ChemiseAddHandler(app))
	r.Get("/beate/add", BeauteAddHandler(app))
	r.Get("/perique/add", PeriqueAddHandler(app))

	/**r.Post("/create/soulier", CreateSoulierHandler(app))
	r.Post("/create/soulier", CreateChemiseHandler(app))
	r.Post("/create/soulier", CreatePeriqueHandler(app))*/
	r.Post("/create/soulier", CreateBeauteHandler(app))
	return r
}

func CreateBeauteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" we are creating Beaute")
		r.ParseForm()
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		reader := bufio.NewReader(file)
		content, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(content)
		ItemName := r.PostFormValue("ItemName")
		size := r.PostFormValue("size")
		color := r.PostFormValue("color")
		description := r.PostFormValue("desc")
		//photo1 :=r.PostFormValue("photo1")

		user := items.BeautyMakeup{ItemName, size, description, color}

		fmt.Println(encoded, " ", user) //result, err := customer.CreateCustomer(user)
		/*if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		app.InfoLog.Println("create response is :", result)
		http.Redirect(w, r, "/", 301)*/
	}
}

func ChemiseAddHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			Entities []itemsIO.ShoesItem
		}

		files := []string{
			app.Path + "itemAdd/chemiseAdd.html",
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

func BeauteAddHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			Entities []itemsIO.ShoesItem
		}

		files := []string{
			app.Path + "itemAdd/beauteAdd.html",
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

func PeriqueAddHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			Entities []itemsIO.ShoesItem
		}

		files := []string{
			app.Path + "itemAdd/soulierAdd.html",
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

func SoulierAddHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			Entities []itemsIO.ShoesItem
		}

		files := []string{
			app.Path + "itemAdd/soulierAdd.html",
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

func SoulierItemHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := itemsIO.GetShoes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []itemsIO.ShoesItem
		}
		data := PageData{resp}

		files := []string{
			app.Path + "items/soulierTable.html",
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

func PeriqueItemHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := itemsIO.GetHairs()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []itemsIO.HairItem
		}
		files := []string{
			app.Path + "items/periqueTable.html",
		}
		data := PageData{resp}

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

func BeauteItemHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := itemsIO.GetBeauties()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []itemsIO.BeautyItem
		}
		data := PageData{resp}

		files := []string{
			app.Path + "items/beauteTable.html",
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

func PantalonHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := itemsIO.GetCloths()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []itemsIO.ClothsItem
		}
		data := PageData{resp}

		files := []string{
			app.Path + "items/pantalonTable.html",
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

func ChemiseItemHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := itemsIO.GetCloths()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []itemsIO.ClothsItem
		}
		data := PageData{resp}
		files := []string{
			app.Path + "items/chemiseTable.html",
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

func indexHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "category.html",
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
