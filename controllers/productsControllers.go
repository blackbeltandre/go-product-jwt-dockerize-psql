package controllers

import (
	"encoding/json"
	"go-product/models"
	u "go-product/utils"
	"net/http"
	"strconv"
	"fmt"
)

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	product.ProductUserID = user
	resp := product.Create()
	u.Respond(w, resp)
}

var GetProduct = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetProducts(id)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetProductDetail = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	ProductUserID := r.Context().Value("user").(uint)
	data := models.GetProductDetail(id,ProductUserID)
	resp := u.Message(true, "Display Detail Product Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var DeleteProduct = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	ProductUserID := r.Context().Value("user").(uint)
	data := models.DeleteProduct(id,ProductUserID)
	resp := u.Message(true, "Delete Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var UpdateProduct = func(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	product := &models.Product{}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	ProductUserID := r.Context().Value("user").(uint)
	data := product.UpdateProduct(id,ProductUserID)
	resp := u.Message(true, "Update Success")
	resp["data"] = data
	u.Respond(w, resp)
}


