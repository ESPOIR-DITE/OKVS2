package item

import (
	"OKVS2/config"
	"OKVS2/domain/gender"
	"OKVS2/domain/items"
	"OKVS2/io/image_oi"
	itemsIO "OKVS2/io/items"
	"OKVS2/io/makeUp"
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
	Name  string
	Class string
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
	r.Get("/types/product", TypesProductTypeHandler(app))
	r.Get("/types/address", TypesAddressHandler(app))

	r.Get("/searchProduct/{resetkey}", ReadProductHandler(app))
	r.Get("/addToCard/{resetkeys}", AddToCardHandler(app))

	/**r.Post("/create/soulier", CreateSoulierHandler(app))
	r.Post("/create/soulier", CreateChemiseHandler(app))
	r.Post("/create/soulier", CreatePeriqueHandler(app))*/
	r.Post("/create/soulier", CreateBeauteHandler(app))
	r.Post("/create/gender", CreateGenderHandler(app))
	r.Post("/create/color", CreateColorHandler(app))
	r.Post("/create/braind", CreateBraindHandler(app))
	r.Post("/create/product", CreateProductTypeHandler(app))
	r.Post("/create/address", CreateAddressHandler(app))

	r.Post("/delete/color", DeleteColorHandler(app))
	r.Post("/delete/gender", DeleteGenderHandler(app))
	r.Post("/delete/braind", DeleteBraindHandler(app))
	r.Post("/delete/product", DeleteProductTypeHandler(app))
	r.Post("/delete/address", DeleteAddressHandler(app))

	r.Post("/search/productType", SearchProductTypeHandler(app))
	r.Post("/search/product", SearchProductHandler(app))

	return r
}

func AddToCardHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		resetKey := chi.URLParam(r, "resetkeys")
		//productTypeId := r.PostFormValue("productpic")
		fmt.Println("product id to add to the card>>>", resetKey, "<<<< User email>>>", userEmail)

		//fmt.Println("product Details to search>>>", productDetails)

		//http.Redirect(w, r, "/", 301)
		return
	}
}

type Numbers struct {
	One, Two, Three int
}

func ReadProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "resetkey")
		//productTypeId := r.PostFormValue("productpic")
		fmt.Println("product id to search>>>", resetKey)
		productDetails, err := makeUp.GetOneItemDetails(resetKey)
		if err != nil {
			if err != nil {
				app.ErrorLog.Println(err.Error())
				http.Redirect(w, r, "/", 301)
				return
			}
		}
		//fmt.Println("product Details to search>>>", productDetails)

		myNumbers := Numbers{0, 1, 3}
		newEnity := items.ViewProduct2{productDetails.ItemId, productDetails.ItemName, productDetails.ItemBrand, productDetails.Price, productDetails.Description, productDetails.Quantity, productDetails.Colors}
		fmt.Println("product Details to search>>>", newEnity)

		type PageData struct {
			Entity  items.ViewProduct2
			Myimage []ImageItems2
			Numbers
		}
		data := PageData{newEnity, GetImageItem2(productDetails.Pictures), myNumbers}
		files := []string{
			app.Path + "items/single-product.html",
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

type ImageItems struct {
	Pic string
}
type ImageItems2 struct {
	Pic      string
	Number   int
	Class    string
	Position string
}

func GetImageItem2(image [][]byte) []ImageItems2 {
	var myList []ImageItems2
	position := [3]string{"First slide", "Second slide", "Third slide"}
	for index, value := range image {
		if index == 0 {
			myList = append(myList, ImageItems2{readImage(value), index, "active", position[index]})
		} else if index != 0 {
			myList = append(myList, ImageItems2{readImage(value), index, "", position[index]})
		}
	}
	return myList
}
func GetImageItem(image []string) []ImageItems {
	entity := []ImageItems{}
	for _, value := range image {
		entity = append(entity, ImageItems{value})
	}
	return entity
}

func SearchProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		productTypeId := r.PostFormValue("productId")
		fmt.Println("product to search>>>", productTypeId)
		var myimages []string
		var colorListe []items.Color
		var sizeListe []items.Size

		productType, _ := types.GetTypes()
		//productId, _ := types.GetProductType(productTypeId)
		product, _ := itemsIO.GetProduct(productTypeId)
		fmt.Println("product product to search>>>", product)

		accounting, _ := itemsIO.GetAccounting(product.Id)
		fmt.Println("product accounting to search>>>", accounting)
		itemColorList, _ := types.GetItemColorList(product.Id)
		fmt.Println("product itemColorList to search>>>", itemColorList)

		for _, itemColor := range itemColorList {
			color, _ := types.GetColor(itemColor.ColorId)
			colorListe = append(colorListe, color)
		}
		fmt.Println("product product to search>>>", colorListe)

		itemBrand, _ := types.GetItemBraind(product.Id)
		fmt.Println("product itemBrand to search>>>", itemBrand)

		braind, _ := types.GetBrand(itemBrand.BraindId)
		fmt.Println("product braind to search>>>", braind)

		itemGender, _ := types.GetItemGender(product.Id)
		fmt.Println("product itemGender to search>>>", itemGender)

		genderdate, _ := types.GetGender(itemGender.GenderId)
		fmt.Println("product genderdate to search>>>", genderdate)

		itemImag, _ := image_oi.GetItemImage(product.Id)
		fmt.Println("product itemImag to search>>>", itemImag)

		productSizes, _ := types.GetPtoductSizeWithItemId(product.Id)
		fmt.Println("product productSizes to search>>>", productSizes)
		for _, itemSize := range productSizes {
			size, _ := types.GetSize(itemSize.SizeId)
			sizeListe = append(sizeListe, size)
		}
		fmt.Println("product sizeListe to search>>>", sizeListe)

		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, _ := image_oi.GetImage(itemImageId.ImageId)
				myimages = append(myimages, readImage(myImage.Image))
			}
		}
		imageStringArry := GetImageItem(myimages)
		fmt.Println("product myimages to search>>>", myimages)

		fmt.Println(" In  product...", product)
		fmt.Println(" In  accounting...", accounting)

		products, _ := itemsIO.GetProducts()
		type PageData struct {
			Product  items.Products
			Account  items.Accounting
			Color    []items.Color
			Braind   items.Braind
			Gender   gender.Gender
			MySize   []items.Size
			Myimage  []ImageItems
			Entities []items.Type
			Products []items.Products
		}
		data := PageData{product, accounting, colorListe, braind, genderdate, sizeListe, imageStringArry, productType, products}
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
func readImage(byteImage []byte) string {
	mybyte := string(byteImage)
	return mybyte
}

func SearchProductTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		productType := []items.Type{}
		product := items.Products{}

		productTypeId := r.PostFormValue("productId")
		fmt.Println("product to search>>>", productTypeId)
		productId, err := types.GetProductType(productTypeId)
		/**var myimages []string
		var colorListe []items.Color
		var sizeListe []items.Size*/

		if err == nil {
			productType, _ = types.GetTypes()
			product, _ := itemsIO.GetProduct(productId.ItemId)
			fmt.Println("product product to search>>>", product)
		}

		//fmt.Println("product product to search>>>", product)

		/**accounting, _ := itemsIO.GetAccounting(product.Id)
				fmt.Println("product accounting to search>>>", accounting)
				itemColorList, _ := types.GetItemColorList(product.Id)
				fmt.Println("product itemColorList to search>>>", itemColorList)


				for _, itemColor := range itemColorList {
					color, _ := types.GetColor(itemColor.ColorId)
					colorListe = append(colorListe, color)
				}
				fmt.Println("product product to search>>>", colorListe)


				itemBrand, _ := types.GetBrand(product.Id)
				fmt.Println("product itemBrand to search>>>", itemBrand)

				braind, _ := types.GetBrand(itemBrand.BraindId)
				fmt.Println("product braind to search>>>", braind)


				itemGender, _ := types.GetItemGender(product.Id)
				fmt.Println("product itemGender to search>>>", itemGender)

				genderdate, _ := types.GetGender(itemGender.GenderId)
				fmt.Println("product genderdate to search>>>", genderdate)


				itemImag, _ := image_oi.GetItemImage(product.Id)
				fmt.Println("product itemImag to search>>>", itemImag)

				productSizes, _ := types.GetPtoductSizeWithItemId(product.Id)
				for _, itemSize := range productSizes {
					size, _ := types.GetSize(itemSize.SizeId)
					sizeListe = append(sizeListe, size)
				}
				fmt.Println("product sizeListe to search>>>", sizeListe)


				if itemImag != nil {
					for _, imageId := range itemImag {
						myImage, _ := image_oi.GetImage(imageId.Id)
						myimages = append(myimages, readImage(myImage.Image))
					}
				}

				fmt.Println(" In  product...", product)
				fmt.Println(" In  accounting...", accounting)
		**/
		type PageData struct {
			Product  items.Products
			Entities []items.Type
			/**Account items.Accounting
			Color   []items.Color
			Braind  items.Braind
			Gender  gender.Gender
			MySize  []items.Size
			Myimage []string*/
		}
		data := PageData{product, productType}
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

func DeleteProductTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  DeleteColorHandler...")

		r.ParseForm()
		ProductdId := r.PostFormValue("ProductdId")
		fmt.Println(" what we are delete ", ProductdId)
		type PageData struct {
			Entities []items.ProductType
		}
		if ProductdId != "" {
			_, nill := types.DeleteProductType(ProductdId)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := types.GetProductTypes()

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

func CreateProductTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  CreateColorHandler...")

		r.ParseForm()
		ProductdName := r.PostFormValue("ProductdName")
		Description := r.PostFormValue("Description")
		fmt.Println(" what we are creating ", ProductdName, " and ", Description)

		type PageData struct {
			Entities []items.ProductType
		}

		if ProductdName != "" && Description != "" {
			_, nill := types.CreateProductType(ProductdName, Description)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := types.GetProductTypes()

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

func TypesProductTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" In  TypesProductTypeHandler...")
		type PageData struct {
			Entities []items.ProductType
		}
		data, nill := types.GetProductTypes()

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
		var res = Results{}

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
		if err != nil && result == false {
			app.ErrorLog.Println(err.Error())
			data = " could not upload the details"
			fmt.Println(err, "  this is the erro")
			res = Results{data, "danger"}
		}
		//app.InfoLog.Println("create response is :", result)
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
		fmt.Println("the read colors>>>", color)
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
		if result != false {
			data = "upload successful"
			res = Results{data, "success"}
		}
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
		res := Results{StringValidatio, ""}
		gender, nill := types.GetGenders()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading gender", nill)
			return
		}
		color, nill := types.GetColors()
		fmt.Println("  reading itemType", color)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading color", nill)
			return
		}
		braind, nill := types.GetBrainds()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading braind", nill)
			return
		}
		size, nill := types.GetSizes()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading size", nill)
			return
		}
		itemType, nill := types.GetTypes()
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
		//productTypes, err := types.GetTypes()
		products, err := itemsIO.GetProducts()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Products []items.Products
		}
		data := PageData{products}
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
