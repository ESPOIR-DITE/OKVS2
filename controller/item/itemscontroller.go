package item

import (
	"OKVS2/config"
	helperUser "OKVS2/controller/users"
	"OKVS2/domain/gender"
	"OKVS2/domain/items"
	"OKVS2/domain/users"
	"OKVS2/io/accountting_io"
	itemsIO "OKVS2/io/item_io"
	"OKVS2/io/item_io/brand"
	"OKVS2/io/item_io/color"
	gender2 "OKVS2/io/item_io/gender"
	"OKVS2/io/item_io/image"
	size2 "OKVS2/io/item_io/size"
	"OKVS2/io/item_io/type"
	"OKVS2/io/joins"
	"OKVS2/io/users_io/address"
	"OKVS2/io/users_io/admin"
	gender3 "OKVS2/io/users_io/gender"
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Results struct {
	Name  string
	Class string
}

//type Gender gender.Gender

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	r.Get("/", indexHanler(app))
	r.Get("/soulier/table", ItemViewHandler(app))
	r.Get("/chemise/table", ChemiseItemHanler(app))
	r.Get("/pantalon/table", PantalonHanler(app))
	r.Get("/beate/table", BeauteItemHanler(app))
	r.Get("/perique/table", PeriqueItemHanler(app))
	r.Get("/item/add", ItemAddHandler(app))
	r.Get("/soulier/add", SoulierAddHandler(app))
	r.Get("/chemise/add", ChemiseAddHandler(app))
	r.Get("/beate/add", BeauteAddHandler(app))
	r.Get("/perique/add", PeriqueAddHandler(app))

	r.Get("/viewProducts/table", ViewProductHandler(app))

	/***TYPES*/
	r.Get("/types", TypesHandler(app))
	r.Get("/types/gender", TypesGenderHandler(app))
	r.Get("/types/color", TypesColorHandler(app))
	r.Get("/types/braind", TypesBraindHandler(app))
	r.Get("/types/product", TypesProductTypeHandler(app))
	r.Get("/types/address", TypesAddressHandler(app))

	//r.Get("/search/product/typeId/{resetkey}", ReadProductTypeIdHandler(app))
	//r.Post("/search/product/type", ReadProductTypeHandler(app))
	/**this link return the product details for a customer to view the product before adding to the card**/
	r.Get("/searchProduct/{resetkey}", ReadProductHandler(app))
	//r.Get("/addToCard/{resetkeys}", AddToCardHandler(app))

	r.Get("/getproducts/{productTypeId}", GetProductsHandler(app))

	/**r.Post("/create/soulier", CreateSoulierHandler(app))
	r.Post("/create/soulier", CreateChemiseHandler(app))
	r.Post("/create/soulier", CreatePeriqueHandler(app))*/

	r.Post("/create/soulier", CreateBeauteHandler(app))
	r.Post("/create/gender", CreateGenderHandler(app))
	r.Post("/create/color", CreateColorHandler(app))
	r.Post("/create/braind", CreateBraindHandler(app))
	r.Post("/create/product", CreateTypeHandler(app))
	r.Post("/create/address", CreateAddressHandler(app))

	r.Post("/delete/color", DeleteColorHandler(app))
	r.Get("/delete/gender/{genderId}", DeleteGenderHandler(app))
	r.Post("/delete/braind", DeleteBraindHandler(app))
	r.Post("/delete/product", DeleteProductTypeHandler(app))
	r.Post("/delete/address", DeleteAddressHandler(app))

	r.Post("/search/productType", SearchProductTypeHandler(app))
	r.Post("/search/product", SearchProductHandler(app))

	r.Post("/update/item/{itemId}", UpdateItemHandler(app))
	r.Post("/update/image", UpdateImageHandler(app))

	r.Post("/update/itemContent", UpdateItemContent(app))

	return r
}

func UpdateItemContent(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		genderId := r.PostFormValue("gender")
		braindId := r.PostFormValue("braind")
		ifColor := r.PostFormValue("ifColor")
		ifSize := r.PostFormValue("ifSize")
		itemName := r.PostFormValue("itemName")
		quantity, _ := strconv.Atoi(r.PostFormValue("quantity"))
		price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
		description := r.PostFormValue("description")
		mysize := r.Form["size"]
		mycolor := r.Form["color"]
		//color := r.PostFormValue("color")
		productId := r.PostFormValue("productId")
		accountId := r.PostFormValue("accountId")

		var productSizeList []items.ItemSize
		var productColorList []items.ItemColor

		fmt.Println(ifColor, " <<<<ifColor||ifSize>>>>>", ifSize)

		fmt.Println(mysize, " <<<<mysize||mycolor>>>>>", mycolor)

		fmt.Println(productId, " <<<<productId||file>>>>>", itemName, "    description>>>>", description)

		//HERE WE CHECK IF THE CHECK BUTTON IS ON WE HAVE TO DELETE ALL THE
		if ifSize == "on" {
			psdeleteResult, err := size2.DeleteAllOfProductSize(productId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), "DeleteAllOfProductSize(productId)")
			} else {
				for _, valeu := range mysize {
					newProductSize := items.ItemSize{"", productId, valeu}
					productSizeList = append(productSizeList, newProductSize)
				}
				resultProductSize, err := size2.CreateAllProductSize(productSizeList)
				if err != nil {
					fmt.Println("error creating productTypeList", err)
				}
				fmt.Println("resultProductSize", resultProductSize)
			}
			fmt.Println(psdeleteResult, " <<<<psdeleteResult")
		}
		if ifColor == "on" {
			pcdeleteReult, err := color.DeleteAllOfItemColor(productId)
			fmt.Println("result for DeleteAllOfItemColor", pcdeleteReult)
			if err != nil {
				fmt.Println("error deleteing products color")
			} else {
				for _, value := range mycolor {
					if value != "" {
						newproductColorList := items.ItemColor{"", productId, value}
						productColorList = append(productColorList, newproductColorList)
					}
				}
				if len(productColorList) != 0 {
					fmt.Println("we are in if len(productColorList)!=0{")
					pcReult, err := color.CreateAllOfItemColors(productColorList)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						fmt.Println("error deleteing products color")
					}
					fmt.Println("result for CreateAllOfItemColors", pcReult)
				}

			}
		}

		if braindId != "" {
			itembraind, err := brand.ReadWithItemId(productId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), " error Getting braind")
			} else {
				newBraind := items.ItemBrand{itembraind.Id, braindId, itembraind.ItemId}
				pbraindResult, err := brand.UpdateItemBrand(newBraind)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					fmt.Println("error deleteing products color")
				}
				fmt.Println("result for pbraindResult", pbraindResult)
			}

		}

		if genderId != "" {
			itemGender, err := gender2.ReadItemGenderWithItemId(productId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), " error Getting braind")
			} else {
				newBraind := items.ItemGender{itemGender.Id, genderId, itemGender.ItemId}
				pGenderResult, err := gender2.UpdateItemGender(newBraind)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					fmt.Println("error deleteing products gender")
				}
				fmt.Println("result for pGenderResult", pGenderResult)
			}
		}

		if itemName != "" || description != "" || productId != "" {
			product := items.Products{productId, itemName, description}
			_, err := itemsIO.UpdateProduct(product)
			if err != nil {
				fmt.Println("error updating products")
			}
		}
		fmt.Println(accountId, " <<<<accountId||quantity>>>>>", quantity, "    price>>>>", price)
		if accountId != "" || quantity != 0 || price != 0 {
			account := items.Accounting{accountId, price, quantity}
			_, err := accountting_io.UpdateAccounting(account)
			if err != nil {
				fmt.Println("error updating account")
			}
		}

		var myimages []string
		var colorListe []items.Color
		var sizeListe []items.Size
		var theImage []image_id

		productType, err := _type.GetTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetTypes()")
		}

		//reading the item that we just edite
		product, err := itemsIO.GetProduct(productId)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetProduct(productId)")
		} //should be an else here like else{}!!!!!

		//fmt.Println("product product to search>>>", product)

		accounting, err := accountting_io.GetAccounting(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetAccounting(product.Id")
		}
		//fmt.Println("product accounting to search>>>", accounting)
		itemColorList, err := color.GetItemColorList(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetAccounting(product.Id")
		} else {
			for _, itemColor := range itemColorList {
				color, err := color.GetColor(itemColor.ColorId)
				if err != nil {
					app.ErrorLog.Println(err.Error(), "erro GetColor(itemColor.ColorId)")
				}
				colorListe = append(colorListe, color)
			}
		}
		//fmt.Println("product itemColorList to search>>>", itemColorList)

		//fmt.Println("product product to search>>>", colorListe)

		itemBrand, err := brand.ReadWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetItemBrand(product.Id)")
		}
		//fmt.Println("product itemBrand to search>>>", itemBrand)

		braind, err := brand.GetBrand(itemBrand.BraindId)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetBrand(itemBrand.BraindId)")
		}
		//fmt.Println("product braind to search>>>", braind)

		itemGender, err := gender2.ReadItemGenderWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro ReadItemGenderWithItemId(product.Id)")
		}
		//fmt.Println("product itemGender to search>>>", itemGender)

		genderdate, err := gender3.GetGender(itemGender.GenderId)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetGender(itemGender.GenderId)")
		}
		//fmt.Println("product genderdate to search>>>", genderdate)

		//reading all the pictures of an item
		itemImag, err := image.GetItemImage(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetItemImage(product.Id)")
		}
		//fmt.Println("product itemImag to search>>>", itemImag)

		productSizes, err := size2.GetPtoductSizeWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetPtoductSizeWithItemId(product.Id)")
		}
		//fmt.Println("product productSizes to search>>>", productSizes)
		for _, itemSize := range productSizes {
			size, err := size2.GetSize(itemSize.SizeId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), "erro GetSize(itemSize.SizeId)")
			}
			sizeListe = append(sizeListe, size)
		}
		//fmt.Println("product sizeListe to search>>>", sizeListe)

		//reading An image
		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, err := image.GetImage(itemImageId.ImageId)
				if err != nil {
					app.ErrorLog.Println(err.Error(), "erro GetImage(itemImageId.ImageId)")
				}
				theImage = append(theImage, image_id{readImage(myImage.Image), myImage.Id})
			}
		}

		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, err := image.GetImage(itemImageId.ImageId)
				if err != nil {
					app.ErrorLog.Println(err.Error(), "erro GetImage(itemImageId.ImageId)")
				}
				myimages = append(myimages, readImage(myImage.Image))
			}
		}
		imageStringArry := GetImageItem(myimages)

		products, err := itemsIO.GetProducts()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetProducts()")
		}

		size, err := size2.GetSizes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading size")
		}
		colors, err := color.GetColors()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading colors")
		}
		genders, err := gender3.GetGenders()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading genders")
		}

		brainds, err := brand.GetBrainds()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading brainds")
		}

		type PageData struct {
			Product    items.Products
			Account    items.Accounting
			Color      []items.Color
			Braind     items.Brand
			Gender     gender.Gender
			MySize     []items.Size
			Myimage    []ImageItems
			Entities   []items.TypeOfItem
			Products   []items.Products
			Thepicture []image_id
			Sizes      []items.Size
			Colors     []items.Color
			Genders    []gender.Gender
			Brainds    []items.Brand
		}
		data := PageData{product, accounting, colorListe, braind, genderdate, sizeListe, imageStringArry, productType, products, theImage, size, colors, genders, brainds}
		files := []string{
			app.Path + "item_io/productsSearchResult.html",
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

func UpdateImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, handler, erreur := r.FormFile("file")
		imageId := r.PostFormValue("Id")

		fmt.Println(imageId, " <<<<imageId||file>>>>>", file)

		if file != nil {
			fmt.Println(" read successful")
			if erreur != nil {
				fmt.Println(erreur, "<<<<<<could not upload the details>>>>>>>", handler)
			}
			reader := bufio.NewReader(file)
			content, _ := ioutil.ReadAll(reader)

			//sliceOfImage := []byte{content}

			myImage := items.Images{imageId, content}

			_, erre := image.UpdateImage(myImage)
			if erre != nil {
				fmt.Println(erre, "<<<<<<could not update picture the details>>>>>>>")
			}
		}

		var myimages []string
		var colorListe []items.Color
		var sizeListe []items.Size
		var braind items.Brand
		var genderdate gender.Gender

		var theImage []image_id

		productType, _ := _type.GetTypes()
		//fetching product Id trough itemImage
		productTypeId, errImage := image.ReadWithImageId(imageId)
		if errImage != nil {
			fmt.Println(errImage, "<<<<<<could not read productTypeId>>>>>>>")
		}
		//productId, _ := types.GetProductType(productTypeId)
		//reading the item that we just edite
		product, _ := itemsIO.GetProduct(productTypeId.ItemId)
		//fmt.Println("product product to search>>>", product)

		accounting, _ := accountting_io.GetAccounting(product.Id)
		//fmt.Println("product accounting to search>>>", accounting)
		itemColorList, _ := color.GetItemColorList(product.Id)
		//fmt.Println("product itemColorList to search>>>", itemColorList)

		for _, itemColor := range itemColorList {
			color, _ := color.GetColor(itemColor.ColorId)
			colorListe = append(colorListe, color)
		}
		//fmt.Println("product product to search>>>", colorListe)

		/****Getting the braind of this item***/
		itemBrand, err := brand.ReadWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro ReadWithItemId(product.Id)")
		} else {
			braind, err = brand.GetBrand(itemBrand.BraindId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), "erro GetBrand(itemBrand.BraindId)")
			}
		}

		//fmt.Println("product braind to search>>>", braind)

		itemGender, err := gender2.GetItemGender(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetItemGender(product.Id)")
		} else {
			genderdate, err = gender3.GetGender(itemGender.GenderId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), "erro GetGender(itemGender.GenderId)")
			}
		}
		//fmt.Println("product itemGender to search>>>", itemGender)

		//fmt.Println("product genderdate to search>>>", genderdate)

		//reding all the pictures of an item
		itemImag, err := image.GetItemImage(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetItemImage(product.Id)")
		}
		//fmt.Println("product itemImag to search>>>", itemImag)

		productSizes, err := size2.GetPtoductSizeWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetPtoductSizeWithItemId(product.Id)")
		}
		//fmt.Println("product productSizes to search>>>", productSizes)
		for _, itemSize := range productSizes {
			size, _ := size2.GetSize(itemSize.SizeId)
			sizeListe = append(sizeListe, size)
		}
		//fmt.Println("product sizeListe to search>>>", sizeListe)

		//reading An image
		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, _ := image.GetImage(itemImageId.ImageId)
				theImage = append(theImage, image_id{readImage(myImage.Image), myImage.Id})
			}
		}

		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, _ := image.GetImage(itemImageId.ImageId)
				myimages = append(myimages, readImage(myImage.Image))
			}
		}
		imageStringArry := GetImageItem(myimages)

		products, _ := itemsIO.GetProducts()

		size, err := size2.GetSizes()
		if err != nil {
			println("erroe reading size")
		}
		colors, err := color.GetColors()
		if err != nil {
			println("erroe reading colors")
		}

		genders, err := gender3.GetGenders()
		if err != nil {
			println("erroe reading colors")
		}

		brainds, err := brand.GetBrainds()
		if err != nil {
			println("erroe reading colors")
		}

		type PageData struct {
			Product    items.Products
			Account    items.Accounting
			Color      []items.Color
			Braind     items.Brand
			Gender     gender.Gender
			MySize     []items.Size
			Myimage    []ImageItems
			Entities   []items.TypeOfItem
			Products   []items.Products
			Thepicture []image_id
			Sizes      []items.Size
			Colors     []items.Color
			Genders    []gender.Gender
			Brainds    []items.Brand
		}
		data := PageData{product, accounting, colorListe, braind, genderdate, sizeListe, imageStringArry, productType, products, theImage, size, colors, genders, brainds}
		files := []string{
			app.Path + "item_io/productsSearchResult.html",
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

func UpdateItemHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productTypeId := chi.URLParam(r, "itemId")
		var myimages []string
		var colorListe []items.Color
		var sizeListe []items.Size

		productType, _ := _type.GetTypes()
		product, _ := itemsIO.GetProduct(productTypeId)
		//fmt.Println("product product to search>>>", product)

		accounting, _ := accountting_io.GetAccounting(product.Id)
		//fmt.Println("product accounting to search>>>", accounting)
		itemColorList, _ := color.GetItemColorList(product.Id)
		//fmt.Println("product itemColorList to search>>>", itemColorList)

		for _, itemColor := range itemColorList {
			color, _ := color.GetColor(itemColor.ColorId)
			colorListe = append(colorListe, color)
		}
		//fmt.Println("product product to search>>>", colorListe)

		itemBrand, _ := brand.GetItemBrand(product.Id)
		//fmt.Println("product itemBrand to search>>>", itemBrand)

		braind, _ := brand.GetBrand(itemBrand.BraindId)
		//fmt.Println("product braind to search>>>", braind)

		itemGender, _ := gender2.GetItemGender(product.Id)
		//fmt.Println("product itemGender to search>>>", itemGender)

		genderdate, _ := gender3.GetGender(itemGender.GenderId)
		//fmt.Println("product genderdate to search>>>", genderdate)

		itemImag, _ := image.GetItemImage(product.Id)
		//fmt.Println("product itemImag to search>>>", itemImag)

		productSizes, _ := size2.GetPtoductSizeWithItemId(product.Id)
		//fmt.Println("product productSizes to search>>>", productSizes)
		for _, itemSize := range productSizes {
			size, _ := size2.GetSize(itemSize.SizeId)
			sizeListe = append(sizeListe, size)
		}
		//fmt.Println("product sizeListe to search>>>", sizeListe)

		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, _ := image.GetImage(itemImageId.ImageId)
				myimages = append(myimages, readImage(myImage.Image))
			}
		}
		imageStringArry := GetImageItem(myimages)
		//fmt.Println("product myimages to search>>>", myimages)

		//fmt.Println(" In  product...", product)
		//fmt.Println(" In  accounting...", accounting)

		products, _ := itemsIO.GetProducts()
		type PageData struct {
			Product  items.Products
			Account  items.Accounting
			Color    []items.Color
			Braind   items.Brand
			Gender   gender.Gender
			MySize   []items.Size
			Myimage  []ImageItems
			Entities []items.TypeOfItem
			Products []items.Products
		}
		data := PageData{product, accounting, colorListe, braind, genderdate, sizeListe, imageStringArry, productType, products}

		files := []string{
			app.Path + "item_io/item_update.html",
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

func GetProductsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productType := chi.URLParam(r, "productTypeId")

		productTypes, err := _type.GetTypes()
		fmt.Println("productTypes>>>", productTypes)
		if err != nil {
			fmt.Println("error reading types")
			app.ErrorLog.Println(err.Error())
		}

		var productList []items.Products
		fmt.Println("productType>>>> ", productType)
		app.Session.Put(r.Context(), "typeId", productType)

		productsTypeObj, err := _type.GetAllOfProductType(productType)
		if err != nil {
			fmt.Println("an error in productTypeObje in GetProductsHandler")
			app.ErrorLog.Println(err.Error())
		}
		for _, value := range productsTypeObj {
			products, err := itemsIO.GetProduct(value.ItemId)
			if err != nil {
				fmt.Println("an error in for loop in GetProductsHandler")
				app.ErrorLog.Println(err.Error())
			}
			productList = append(productList, products)
		}
		fmt.Println("productsList>>>", productList)

		type PageData struct {
			Products     []items.Products
			ProductTypes []items.TypeOfItem
		}
		data := PageData{productList, productTypes}

		files := []string{
			app.Path + "item_io/itemProduct.html",
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
		//render.JSON(w, r, productList)
	}
}
func ReadProductTypeIdHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productType := chi.URLParam(r, "resetkey")

		fmt.Println("productType>>>> ", productType)
		app.Session.Put(r.Context(), "typeId", productType)

		//productTypeObje, err := types.GetProductType(productType)

		http.Redirect(w, r, "/item/search/product/type", 301)
		return
	}
}

func ReadProductTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//productType := chi.URLParam(r, "resetkey")
		//productType := app.Session.GetString(r.Context(), "typeId")
		//fmt.Println("productType>>>>", productType)
		//
		//productTypeObje, err := types.GetProductType(productType)
		//if err != nil {
		//	fmt.Println("an error in productTypeObje in ReadProductTypeHandler")
		//	app.ErrorLog.Println(err.Error())
		//}

		r.ParseForm()
		productTypeId := r.PostFormValue("productId")

		var myimages []string
		var colorListe []items.Color
		var sizeListe []items.Size

		var product items.Products
		productTypes, _ := _type.GetTypes()
		//productId, _ := types.GetProductType(productTypeId)
		if productTypeId != "" {
			product, _ = itemsIO.GetProduct(productTypeId)

		}
		//fmt.Println("product product to search>>>", product)

		accounting, _ := accountting_io.GetAccounting(product.Id)
		//fmt.Println("product accounting to search>>>", accounting)
		itemColorList, _ := color.GetItemColorList(product.Id)
		//fmt.Println("product itemColorList to search>>>", itemColorList)

		for _, itemColor := range itemColorList {
			color, _ := color.GetColor(itemColor.ColorId)
			colorListe = append(colorListe, color)
		}
		//fmt.Println("product product to search>>>", colorListe)

		itemBrand, _ := brand.GetItemBrand(product.Id)
		//fmt.Println("product itemBrand to search>>>", itemBrand)

		braind, _ := brand.GetBrand(itemBrand.BraindId)
		//fmt.Println("product braind to search>>>", braind)

		itemGender, _ := gender2.GetItemGender(product.Id)
		//fmt.Println("product itemGender to search>>>", itemGender)

		genderdate, _ := gender3.GetGender(itemGender.GenderId)
		fmt.Println("product genderdate to search>>>", genderdate)

		itemImag, _ := image.GetItemImage(product.Id)
		//fmt.Println("product itemImag to search>>>", itemImag)

		productSizes, _ := size2.GetPtoductSizeWithItemId(product.Id)
		//fmt.Println("product productSizes to search>>>", productSizes)
		for _, itemSize := range productSizes {
			size, _ := size2.GetSize(itemSize.SizeId)
			sizeListe = append(sizeListe, size)
		}
		//fmt.Println("product sizeListe to search>>>", sizeListe)

		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, _ := image.GetImage(itemImageId.ImageId)
				myimages = append(myimages, readImage(myImage.Image))
			}
		}
		imageStringArry := GetImageItem(myimages)
		//fmt.Println("product myimages to search>>>", myimages)

		//fmt.Println(" In  product...", product)
		//fmt.Println(" In  accounting...", accounting)

		products, _ := itemsIO.GetProducts()
		type PageData struct {
			Product  items.Products
			Account  items.Accounting
			Color    []items.Color
			Braind   items.Brand
			Gender   gender.Gender
			MySize   []items.Size
			Myimage  []ImageItems
			Entities []items.TypeOfItem
			Products []items.Products
		}
		data := PageData{product, accounting, colorListe, braind, genderdate, sizeListe, imageStringArry, productTypes, products}
		files := []string{
			app.Path + "/item_io/productsSearchResult.html",
			app.Path + "/template/admin_navbar.html",
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

func ViewProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			Entity       []items.ViewItem
			ProductTypes []items.TypeOfItem
		}
		products, err := joins.ViewAllItems()
		//fmt.Println(products)
		if err != nil {
			fmt.Println("an error in ViewProductHandler in itemsController")
			app.ErrorLog.Println(err.Error())
		}
		productType, err := _type.GetTypes()
		if err != nil {
			fmt.Println("an error in ViewProductHandler when reading productType")
			app.ErrorLog.Println(err.Error())
		}
		data := PageData{products, productType}
		files := []string{
			app.Path + "item_io/ProductTable.html",
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

//this method help to set the single product page of the selected item
func ReadProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var message string
		var class string
		var Manager = false
		var user users.Customer

		resetKey := chi.URLParam(r, "resetkey")
		//productTypeId := r.PostFormValue("productpic")
		//fmt.Println("product id to search>>>", resetKey)
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		message, class, Manager, user = helperUser.GetUserDetails(userEmail)
		if Manager == true {
			_, err := admin.GetAdmin(userEmail)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				Manager = true
				http.Redirect(w, r, "/user/managementwelcom", 301)
				return
			}
		}

		productDetails, err := joins.GetOneItemDetails(resetKey)
		//fmt.Println("productDetails>>>", productDetails)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/", 301)
			return
		}
		//fmt.Println("product Details to search>>>", productDetails)

		myNumbers := Numbers{0, 1, 3}
		newEnity := items.ViewItem2{productDetails.ItemId, productDetails.ItemName, productDetails.ItemBrand, productDetails.Price, productDetails.Description, productDetails.Quantity, productDetails.Colors}
		//fmt.Println("product Details to search>>>", newEnity)

		type PageData struct {
			EntityProduct items.ViewItem2
			Myimage       []ImageItems2
			Numbers
			Entity CardeData
			MyUser
			Manager bool
			User    users.Customer
		}
		data1 := CardeData{message, class}
		data := PageData{newEnity, GetImageItem2(productDetails.Pictures), myNumbers, data1, MyUser{userEmail}, Manager, user}
		files := []string{
			app.Path + "item_io/single-product.html",
			app.Path + "customer-template/toolbarTemplate.html",
			app.Path + "customer-template/navbar.html",
			app.Path + "customer-template/reviewTemplate.html",
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

//im creating this class so that i can get an image with it id
type image_id struct {
	Image string
	Id    string
}

func SearchProductHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		productId := r.PostFormValue("productId")
		//fmt.Println("product to search>>>", productTypeId)
		var myimages []string
		var colorListe []items.Color
		var sizeListe []items.Size
		var theImage []image_id
		var braind items.Brand
		var genderdate gender.Gender

		productType, err := _type.GetTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetTypes()")
		}
		//productId, _ := types.GetProductType(productTypeId)
		product, err := itemsIO.GetProduct(productId)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetProduct(productId)")
		}
		//fmt.Println("product product to search>>>", product)

		accounting, err := accountting_io.GetAccounting(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetAccounting(product.Id)")
		}
		//fmt.Println("product accounting to search>>>", accounting)
		itemColorList, err := color.GetItemColorList(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetItemColorList(product.Id)")
		}
		//fmt.Println("product itemColorList to search>>>", itemColorList)

		for _, itemColor := range itemColorList {
			color, _ := color.GetColor(itemColor.ColorId)
			colorListe = append(colorListe, color)
		}
		//fmt.Println("product product to search>>>", colorListe)

		itemBrand, err := brand.ReadWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro ReadWithItemId(product.Id)")
		} else {
			braind, err = brand.GetBrand(itemBrand.BraindId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), "erro GetBrand(itemBrand.BraindId)")
			}
		}

		//fmt.Println("product braind to search>>>", braind)

		itemGender, err := gender2.ReadItemGenderWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetBrand(itemBrand.BraindId)")
		} else {
			genderdate, err = gender3.GetGender(itemGender.GenderId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), "erro GetGender(itemGender.GenderId)")
			}
		}
		//fmt.Println("product itemGender to search>>>", itemGender)

		//fmt.Println("product genderdate to search>>>", genderdate)

		//reding all the pictures of an item
		itemImag, err := image.GetItemImage(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetItemImage(product.Id)")
		}
		//fmt.Println("product itemImag to search>>>", itemImag)

		productSizes, err := size2.GetPtoductSizeWithItemId(product.Id)
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erro GetItemImage(product.Id)")
		}
		//fmt.Println("product productSizes to search>>>", productSizes)
		for _, itemSize := range productSizes {
			size, err := size2.GetSize(itemSize.SizeId)
			if err != nil {
				app.ErrorLog.Println(err.Error(), "GetSize(itemSize.SizeId)")
			}
			sizeListe = append(sizeListe, size)
		}
		//fmt.Println("product sizeListe to search>>>", sizeListe)

		//reading An image
		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, _ := image.GetImage(itemImageId.ImageId)
				theImage = append(theImage, image_id{readImage(myImage.Image), myImage.Id})
			}
		}

		if itemImag != nil {
			for _, itemImageId := range itemImag {
				myImage, _ := image.GetImage(itemImageId.ImageId)
				myimages = append(myimages, readImage(myImage.Image))
			}
		}
		imageStringArry := GetImageItem(myimages)

		products, _ := itemsIO.GetProducts()

		size, err := size2.GetSizes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading size")
		}
		colors, err := color.GetColors()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading colors")
		}
		genders, err := gender3.GetGenders()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading colors")
		}

		brainds, err := brand.GetBrainds()
		if err != nil {
			app.ErrorLog.Println(err.Error(), "erroe reading colors")
		}

		type PageData struct {
			Product    items.Products
			Account    items.Accounting
			Color      []items.Color
			Braind     items.Brand
			Gender     gender.Gender
			MySize     []items.Size
			Myimage    []ImageItems
			Entities   []items.TypeOfItem
			Products   []items.Products
			Thepicture []image_id
			Sizes      []items.Size
			Colors     []items.Color
			Genders    []gender.Gender
			Brainds    []items.Brand
		}
		data := PageData{product, accounting, colorListe, braind, genderdate, sizeListe, imageStringArry, productType, products, theImage, size, colors, genders, brainds}
		files := []string{
			app.Path + "item_io/productsSearchResult.html",
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
func readImage(byteImage []byte) string {
	mybyte := string(byteImage)
	return mybyte
}

func SearchProductTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		productType := []items.TypeOfItem{}
		product := items.Products{}

		productTypeId := r.PostFormValue("productId")
		fmt.Println("product to search>>>", productTypeId)
		productId, err := _type.GetProductType(productTypeId)
		/**var myimages []string
		var colorListe []item_io.Color
		var sizeListe []item_io.Size*/

		if err == nil {
			productType, _ = _type.GetTypes()
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


				itemImag, _ := image.GetItemImage(product.Id)
				fmt.Println("product itemImag to search>>>", itemImag)

				productSizes, _ := types.GetPtoductSizeWithItemId(product.Id)
				for _, itemSize := range productSizes {
					size, _ := types.GetSize(itemSize.SizeId)
					sizeListe = append(sizeListe, size)
				}
				fmt.Println("product sizeListe to search>>>", sizeListe)


				if itemImag != nil {
					for _, imageId := range itemImag {
						myImage, _ := image.GetImage(imageId.Id)
						myimages = append(myimages, readImage(myImage.Images))
					}
				}

				fmt.Println(" In  product...", product)
				fmt.Println(" In  accounting...", accounting)
		**/
		type PageData struct {
			Product  items.Products
			Entities []items.TypeOfItem
			/**Account item_io.Accounting
			Color   []item_io.Color
			Brand  item_io.Brand
			Gender  gender.Gender
			MySize  []item_io.Size
			Myimage []string*/
		}
		data := PageData{product, productType}
		files := []string{
			app.Path + "item_io/itemProduct.html",
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
			app.Path + "template/admin_navbar.html",
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
			Entities []users_io.AddressType
		}
		if addressId != "" {
			_, nill := users_io.DeleteAddressType(addressId)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := users_io.GetAddressTypes()

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
			Entities []users_io.AddressType
		}

		if addressdName != "" {
			_, nill := users_io.CreateAddressType(addressdName)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := users_io.GetAddressTypes()

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
			Entities []users_io.AddressType
		}
		data, nill := users_io.GetAddressTypes()

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
			Entities []items.TypeOfItem
		}
		if ProductdId != "" {
			_, nill := _type.DeleteType(ProductdId)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := _type.GetTypes()

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

func CreateTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" In  CreateColorHandler...")

		r.ParseForm()
		typeName := r.PostFormValue("ProductdName")
		//Description := r.PostFormValue("Description")
		fmt.Println(" what we are creating ", typeName)

		type PageData struct {
			Entities []items.TypeOfItem
		}

		if typeName != "" {
			//_, nill := types.CreateProductType(ProductdName, Description)
			_, nill := _type.CreateType(typeName)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")
			}
		}
		data2, err := _type.GetTypes()

		if err != nil {
			app.ErrorLog.Println(err.Error())
			fmt.Println(" Error when reading GetTypes ", data2)
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
			Entities []items.TypeOfItem
		}
		itemType, nill := _type.GetTypes()

		if nill != nil {
			app.ErrorLog.Println(nill.Error())
		}
		Data := PageData{itemType}
		fmt.Println(" we are calling product page", itemType)
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
			Entities []items.Brand
		}
		if braind != "" {
			_, nill := brand.DeleteBraind(braind)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := brand.GetBrainds()

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
			Entities []items.Brand
		}

		if braind != "" {
			_, nill := brand.CreateBraind(braind)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := brand.GetBrainds()

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
			Entities []items.Brand
		}
		data, nill := brand.GetBrainds()

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
			_, nill := color.CreateColors(color)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
				fmt.Println(" Error when creating ")

			}
		}
		data2, err := color.GetColors()

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
			_, nill := color.DeleteColor(color)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := color.GetColors()

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
		data, nill := color.GetColors()

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

		genderName := chi.URLParam(r, "genderId")
		fmt.Println(" what we are delete ", genderName)
		type PageData struct {
			Entities []gender.Gender
		}
		if genderName != "" {
			_, nill := gender3.DeleteGender(genderName)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())
			}
		}
		data2, err := gender3.GetGenders()

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
			_, nill := gender3.CreateGender(genderName)

			if nill != nil {
				app.ErrorLog.Println(nill.Error())

			}
		}
		data2, err := gender3.GetGenders()

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
		data, nill := gender3.GetGenders()

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

		//converting the file into an array of byte
		sliceOfImage := [][]byte{content, content1, content2}
		//a:=item_io.MyImages{content,content1,content2}

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
			app.Path + "template/admin_navbar.html",
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
			ItemTypeData []items.TypeOfItem
			BraindData   []items.Brand
			Result       Results
		}

		mygender, nill := gender3.GetGenders()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		color, nill := color.GetColors()
		fmt.Println("the read colors>>>", color)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		mybraind, nill := brand.GetBrainds()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		mysize, nill := size2.GetSizes()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		myitemType, nill := _type.GetTypes()
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			return
		}
		if result != false {
			data = "Upload Successfully"
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
			ItemTypeData []items.TypeOfItem
			BraindData   []items.Brand
			Result       Results
		}
		res := Results{StringValidatio, ""}
		gender, nill := gender3.GetGenders()
		fmt.Println("  reading gender", gender)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading gender", nill)
			return
		}
		color, nill := color.GetColors()
		fmt.Println("  reading color", color)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading color", nill)
			return
		}
		braind, nill := brand.GetBrainds()
		fmt.Println("  reading braind", braind)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading braind", nill)
			return
		}
		size, nill := size2.GetSizes()
		fmt.Println("  reading size", size)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading size", nill)
			return
		}
		itemType, nill := _type.GetTypes()
		fmt.Println("  reading itemType", itemType)
		if nill != nil {
			app.ErrorLog.Println(nill.Error())
			fmt.Println(" Error reading itemType", nill)
			return
		}
		data := PageData{gender, size, color, itemType, braind, res}
		files := []string{
			app.Path + "itemAdd/addItem.html",
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

func ItemViewHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//productTypes, err := types.GetTypes()
		var products []items.Products

		productTypes, err := _type.GetTypes()
		fmt.Println("productTypes>>>", productTypes)
		if err != nil {
			fmt.Println("error reading types")
			app.ErrorLog.Println(err.Error())
		}
		type PageData struct {
			Products     []items.Products
			ProductTypes []items.TypeOfItem
		}
		data := PageData{products, productTypes}
		files := []string{
			app.Path + "item_io/itemProduct.html",
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
			app.Path + "item_io/periqueTable.html",
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
			app.Path + "item_io/beauteTable.html",
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
			app.Path + "item_io/pantalonTable.html",
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
			app.Path + "item_io/chemiseTable.html",
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

type CardeData struct {
	Mesage string
	Class  string
}
type MyUser struct {
	User string
}

func indexHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := app.Session.GetString(r.Context(), "userEmail")
		message, class, Manager, user := helperUser.GetUserDetails(userEmail)

		type PageData struct {
			Entity CardeData
			MyUser
			Manager bool
			User    users.Customer
		}
		data := PageData{CardeData{message, class}, MyUser{userEmail}, Manager, user}

		files := []string{
			app.Path + "category.html",
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
