package item

import (
	"OKVS2/config"
	"OKVS2/domain/gender"
	"OKVS2/domain/items"
	"OKVS2/io/image_oi"
	itemsIO "OKVS2/io/items"
	"OKVS2/io/types"
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

/**
type Soulier items.Shoes
type Perique items.Hair
type Items items.Items
type ItemSold items.ItemSold
type Cloths items.Cloths
type Beate items.BeautyMakeup*/
type Results struct {
	Name string
}

//type Gender gender.Gender

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	r.Get("/", indexHanler(app))
	r.Get("/soulier/table", SoulierItemHanler(app))
	r.Get("/chemise/table", ChemiseItemHanler(app))
	r.Get("/pantalon/table", PantalonHanler(app))
	r.Get("/beate/table", BeauteItemHanler(app))
	r.Get("/perique/table", PeriqueItemHanler(app))

	r.Get("/item/add", ItemAddHandler(app))

	r.Get("/soulier/add", SoulierAddHandler(app))
	r.Get("/chemise/add", ChemiseAddHandler(app))
	r.Get("/beate/add", BeauteAddHandler(app))
	r.Get("/perique/add", PeriqueAddHandler(app))

	/***TYPES*/
	r.Get("/types", TypesHandler(app))
	r.Get("/types/gender", TypesGenderHandler(app))
	r.Get("/types/color", TypesColorHandler(app))
	r.Get("/types/braind", TypesBraindHandler(app))
	r.Get("/types/product", TypesProductHandler(app))
	r.Get("/types/address", TypesAddressHandler(app))

	/**r.Post("/create/soulier", CreateSoulierHandler(app))
	r.Post("/create/soulier", CreateChemiseHandler(app))
	r.Post("/create/soulier", CreatePeriqueHandler(app))*/
	r.Post("/create/soulier", CreateBeauteHandler(app))
	r.Post("/create/gender", CreateGenderHandler(app))
	r.Post("/create/color", CreateColorHandler(app))
	r.Post("/create/braind", CreateBraindHandler(app))
	r.Post("/create/product", CreateProductHandler(app))
	r.Post("/create/address", CreateAddressHandler(app))

	r.Post("/delete/color", DeleteColorHandler(app))
	r.Post("/delete/gender", DeleteGenderHandler(app))
	r.Post("/delete/braind", DeleteBraindHandler(app))
	r.Post("/delete/product", DeleteProductHandler(app))
	r.Post("/delete/address", DeleteAddressHandler(app))

	r.Post("/search/product", SearchProductHandler(app))

	return r
}
func readImage(byteImage []byte) string {
	mybyte := string(byteImage)
	return mybyte
}
func SearchProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		productId := r.PostFormValue("productId")

		var myimages []string
		product, _ := types.GetProduct(productId)
		accounting, _ := itemsIO.GetAccounting(productId)
		color, _ := types.GetColor(productId)
		braind, _ := types.GetBrand(productId)
		genderdate, _ := types.GetGender(productId)
		itemImag, _ := image_oi.GetItemImage(productId)
		theSize, _ := types.GetSize(productId)

		if itemImag != nil {
			for _, imageId := range itemImag {
				myImage, _ := image_oi.GetImage(imageId.Id)
				myimages = append(myimages, readImage(myImage.Image))
			}
		}

		fmt.Println(" In  product...", product)
		fmt.Println(" In  accounting...", accounting)

		type PageData struct {
			Product items.Products
			Account items.Accounting
			Color   items.Color
			Braind  items.Braind
			Gender  gender.Gender
			MySize  items.Size
			Myimage []string
		}
		data := PageData{product, accounting, color, braind, genderdate, theSize, myimages}
		files := []string{
			app.Path + "items/productsSearchResult.html",
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

func ItemAddHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			Entities []itemsIO.ShoesItem
		}

		files := []string{
			app.Path + "itemAdd/addItem.html",
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

func DeleteAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  DeleteAddressHandler...")

		r.ParseForm()
		addressId := r.PostFormValue("addressId")
		fmt.Println(" what we are delete ", addressId)
		type PageData struct {
			Entities []types.AddressType
		}
		if addressId != "" {
			_, nill := types.DeleteAddressType(addressId)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := types.GetAddressTypes()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", addressId)
		}

		Data := PageData{data2}
		fmt.Println(" we are reading", Data)
		files := []string{
			app.Path + "create_types/addresses_type.html",
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

func CreateAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  CreateAddressHandler...")

		r.ParseForm()
		addressdName := r.PostFormValue("addressdName")

		fmt.Println(" what we are creating ", addressdName)

		type PageData struct {
			Entities []types.AddressType
		}

		if addressdName != "" {
			_, nill := types.CreateAddressType(addressdName)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := types.GetAddressTypes()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", addressdName)

		}

		Data := PageData{data2}
		fmt.Println(" we are creating colore", Data)
		files := []string{
			app.Path + "create_types/addresses_type.html",
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

func TypesAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" In  TypesAddressHandler...")
		type PageData struct {
			Entities []types.AddressType
		}
		data, nill := types.GetAddressTypes()

		if nill != nil {
			app.ErrorLog.Println(nill.Error())
		}
		Data := PageData{data}
		fmt.Println(" we are calling addressType page", Data)
		files := []string{
			app.Path + "create_types/addresses_type.html",
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

func DeleteProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  DeleteColorHandler...")

		r.ParseForm()
		ProductdId := r.PostFormValue("ProductdId")
		fmt.Println(" what we are delete ", ProductdId)
		type PageData struct {
			Entities []items.Products
		}
		if ProductdId != "" {
			_, nill := types.DeleteProduct(ProductdId)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := types.GetProducts()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", ProductdId)
		}

		Data := PageData{data2}
		fmt.Println(" we are reading", Data)
		files := []string{
			app.Path + "create_types/products_type.html",
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

func CreateProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  CreateColorHandler...")

		r.ParseForm()
		ProductdName := r.PostFormValue("ProductdName")
		Description := r.PostFormValue("Description")
		fmt.Println(" what we are creating ", ProductdName, " and ", Description)

		type PageData struct {
			Entities []items.Products
		}

		if ProductdName != "" && Description != "" {
			_, nill := types.CreateProduct(ProductdName, Description)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := types.GetProducts()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", ProductdName)

		}

		Data := PageData{data2}
		fmt.Println(" we are creating colore", Data)
		files := []string{
			app.Path + "create_types/products_type.html",
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

func TypesProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" In  TypesProductHandler...")
		type PageData struct {
			Entities []items.Products
		}
		data, nill := types.GetProducts()

		if nill != nil {
			app.ErrorLog.Println(nill.Error())
		}
		Data := PageData{data}
		fmt.Println(" we are calling product page", Data)
		files := []string{
			app.Path + "create_types/products_type.html",
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

func DeleteBraindHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  DeleteColorHandler...")

		r.ParseForm()
		braind := r.PostFormValue("BraindName")
		fmt.Println(" what we are delete ", braind)
		type PageData struct {
			Entities []items.Braind
		}
		if braind != "" {
			_, nill := types.DeleteBraind(braind)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := types.GetBrainds()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", braind)
		}

		Data := PageData{data2}
		fmt.Println(" we are reading", Data)
		files := []string{
			app.Path + "create_types/brainds_type.html",
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

func CreateBraindHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  CreateColorHandler...")

		r.ParseForm()
		braind := r.PostFormValue("BraindName")
		fmt.Println(" what we are creating ", braind)

		type PageData struct {
			Entities []items.Braind
		}

		if braind != "" {
			_, nill := types.CreateBraind(braind)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := types.GetBrainds()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", braind)

		}

		Data := PageData{data2}
		fmt.Println(" we are creating colore", Data)
		files := []string{
			app.Path + "create_types/brainds_type.html",
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

func TypesBraindHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" In  TypesColorHandler...")
		type PageData struct {
			Entities []items.Braind
		}
		data, nill := types.GetBrainds()

		if nill != nil {
			app.ErrorLog.Println(nill.Error())
		}
		Data := PageData{data}
		fmt.Println(" we are calling color page", Data)
		files := []string{
			app.Path + "create_types/brainds_type.html",
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

func CreateColorHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  CreateColorHandler...")

		r.ParseForm()
		color := r.PostFormValue("ColorName")
		fmt.Println(" what we are creating ", color)

		type PageData struct {
			Entities []items.Color
		}

		if color != "" {
			_, nill := types.CreateColors(color)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := types.GetColors()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", color)

		}

		Data := PageData{data2}
		fmt.Println(" we are creating colore", Data)
		files := []string{
			app.Path + "create_types/colors_type.html",
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

func DeleteColorHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  DeleteColorHandler...")

		r.ParseForm()
		color := r.PostFormValue("ColorName")
		fmt.Println(" what we are delete ", color)
		type PageData struct {
			Entities []items.Color
		}
		if color != "" {
			_, nill := types.DeleteColor(color)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := types.GetColors()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", color)
		}

		Data := PageData{data2}
		fmt.Println(" we are reading", Data)
		files := []string{
			app.Path + "create_types/colors_type.html",
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

func TypesColorHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" In  TypesColorHandler...")
		type PageData struct {
			Entities []items.Color
		}
		data, nill := types.GetColors()

		if nill != nil {
			app.ErrorLog.Println(nill.Error())
		}
		Data := PageData{data}
		fmt.Println(" we are calling color page", Data)
		files := []string{
			app.Path + "create_types/colors_type.html",
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

func DeleteGenderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  DeleteGenderHandler...")

		r.ParseForm()
		genderName := r.PostFormValue("GenderName")
		fmt.Println(" what we are delete ", genderName)
		type PageData struct {
			Entities []gender.Gender
		}
		if genderName != "" {
			_, nill := types.DeleteGender(genderName)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := types.GetGenders()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", genderName)
		}

		Data := PageData{data2}
		fmt.Println(" we are creating Beaute", Data)
		files := []string{
			app.Path + "create_types/geder_type.html",
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

func CreateGenderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  CreateGenderHandler...")

		//var stat string
		r.ParseForm()
		genderName := r.PostFormValue("GenderName")
		fmt.Println(" what we are creating ", genderName)

		type PageData struct {
			Entities []gender.Gender
		}

		if genderName != "" {
			_, nill := types.CreateGender(genderName)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())

			}
		}
		data2, err := types.GetGenders()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading ", genderName)

		}

		Data := PageData{data2}
		fmt.Println(" we are creating Beaute", Data)
		files := []string{
			app.Path + "create_types/geder_type.html",
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

func TypesGenderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  TypesGenderHandler...")

		type PageData struct {
			Entities []gender.Gender
		}
		data, nill := types.GetGenders()

		if nill != nil {
			app.ErrorLog.Println(nill.Error())
		}
		Data := PageData{data}
		fmt.Println(" we are creating Beaute", Data)
		files := []string{
			app.Path + "create_types/geder_type.html",
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

func TypesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  TypesHandler...")

		type PageData struct {
			Entities []itemsIO.ShoesItem
		}

		files := []string{
			app.Path + "create_types/types.html",
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

func CreateBeauteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" we are creating Beaute")
		r.ParseForm()
		fmt.Println(" reading the file")
		file, handler, err := r.FormFile("file")
		file1, handler, err := r.FormFile("file1")
		file2, handler, err := r.FormFile("file2")
		fmt.Println(" read successful")
		var data string

		fmt.Println("********")
		if err != nil {
			fmt.Println(err, "<<<<<<>>>>>>>")
			data = " could not upload the details"
		}
		//defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)
		fmt.Println(" converting to byte array", handler)
		reader := bufio.NewReader(file)
		reader1 := bufio.NewReader(file1)
		reader2 := bufio.NewReader(file2)

		content, _ := ioutil.ReadAll(reader)
		content1, _ := ioutil.ReadAll(reader1)
		content2, _ := ioutil.ReadAll(reader2)
		sliceOfImage := [][]byte{content, content1, content2}
		//a:=items.MyImages{content,content1,content2}

		fmt.Println("converting to byte array successful")
		//encoded := base64.StdEncoding.EncodeToString(content)
		ItemName := r.PostFormValue("ItemName")
		//color := r.PostFormValue("color")
		description := r.PostFormValue("decription")
		itemType := r.PostFormValue("itemType")
		fmt.Println("item type>>>", itemType)
		genders := r.Form["gender"]
		quantity, _ := strconv.Atoi(r.PostFormValue("quantity"))
		price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
		braind := r.PostFormValue("braind")

		fmt.Println("creating an Beauty object")

		size := r.Form["size"]
		Z := r.Form["colors"]
		fmt.Println("item type>>>", Z)

		B := items.MyItemHelper{ItemName, contains(size), description, genders, itemType, quantity, price, sliceOfImage, Z, braind}

		//fmt.Println("creating an Beauty object successful>>>> :",B)

		fmt.Println("sending to backend")

		result, err := itemsIO.CreatBeatyHelper(B)
		fmt.Println("sending to backend successful")

		//fmt.Println(encoded, " ", B) //
		if err != nil {
			app.ErrorLog.Println(err.Error())
			data = " could not upload the details"
			fmt.Println(err, "  this is the erro")
		}
		app.InfoLog.Println("create response is :", result)
		//http.Redirect(w, r, "/", 301)

		files := []string{
			app.Path + "itemAdd/addItem.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		if result == true {
			data = "upload successful"
		}

		type PageData struct {
			GenderData   []gender.Gender
			SizeData     []items.Size
			ColorData    []items.Color
			ItemTypeData []items.Type
			BraindData   []items.Braind
			Result       Results
		}

		mygender, nill := types.GetGenders()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		color, nill := types.GetColors()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		mybraind, nill := types.GetBrainds()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		mysize, nill := types.GetSizes()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		myitemType, nill := types.GetTypes()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		data = "upload successful"
		res := Results{data}
		datatypes := PageData{mygender, mysize, color, myitemType, mybraind, res}

		err = ts.Execute(w, datatypes)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}
func containsImages(slice [][]byte) [][]byte {
	set := [][]byte{}
	for _, s := range slice {
		set = append(slice, s)
	}
	return set
}
func contains(slice []string) []string {
	set := []string{}
	for _, s := range slice {
		set = append(slice, s)
	}
	return set
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

		StringValidatio := ""
		type PageData struct {
			GenderData   []gender.Gender
			SizeData     []items.Size
			ColorData    []items.Color
			ItemTypeData []items.Type
			BraindData   []items.Braind
			Result       Results
		}
		res := Results{StringValidatio}
		gender, nill := types.GetGenders()
		fmt.Println("  reading gender", gender)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading gender", nill)
			return
		}
		color, nill := types.GetColors()
		fmt.Println("  reading color", color)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading color", nill)
			return
		}
		braind, nill := types.GetBrainds()
		fmt.Println("  reading braind", braind)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading braind", nill)
			return
		}
		size, nill := types.GetSizes()
		fmt.Println("  reading size", size)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading size", nill)
			return
		}
		itemType, nill := types.GetTypes()
		fmt.Println("  reading itemType", itemType)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading itemType", nill)
			return
		}
		data := PageData{gender, size, color, itemType, braind, res}
		files := []string{
			app.Path + "itemAdd/addItem.html",
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

func SoulierItemHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := types.GetProducts()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Entities []items.Products
		}
		data := PageData{resp}
		files := []string{
			app.Path + "items/itemProduct.html",
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
