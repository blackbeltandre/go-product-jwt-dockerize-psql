package models

import (
	"fmt"
	"gorm.io/gorm"
	u "go-product/utils"
)

type Product struct {
	gorm.Model
	ProductName   string `json:"product_name"`
	ProductQuantity  string `json:"product_quantity"`
	ProductPrice  string `json:"product_price"`
	ProductImage  string `json:"product_image"`
	ProductUserID  uint `json:"product_user_id"`
	
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (product *Product) Validate() (map[string]interface{}, bool) {
	if product.ProductName == "" {
		return u.Message(false, "Product Name should be on the payload"), false
	}
	if product.ProductQuantity == "" {
		return u.Message(false, "Product Quantity should be on the payload"), false
	}
	if product.ProductPrice == "" {
		return u.Message(false, "Product Price number should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (product *Product) Create() (map[string]interface{}) {

	if resp, ok := product.Validate(); !ok {
		return resp
	}

	GetDB().Create(product)

	resp := u.Message(true, "success")
	resp["product"] = product
	return resp
}

func GetProduct(id uint) (*Product) {

	product := &Product{}
	err := GetDB().Table("products").Where("id = ?", id).First(product).Error
	if err != nil {
		return nil
	}
	return product
}

func GetProducts(ProductUserID uint) ([]*Product) {

	products := make([]*Product, 0)
	err := GetDB().Table("products").Where("product_user_id = ?", ProductUserID).Find(&products).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return products
}

func GetProductDetail(id int, ProductUserID uint) ([]*Product) {
	products := make([]*Product, 0)
	err := GetDB().Where("id = ? and product_user_id = ?", id, ProductUserID).Find(&products).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return products
}
func DeleteProduct(id int, ProductUserID uint) ([]*Product) {
	products := make([]*Product, 0)
	err := GetDB().Where("id = ? and product_user_id = ?", id, ProductUserID).Delete(&products).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return products
}
func (product *Product) UpdateProduct(id int, ProductUserID uint) (map[string]interface{}) {

	if resp, ok := product.Validate(); !ok {
		return resp
	}
	err := GetDB().Where("id = ? and product_user_id = ?", id, ProductUserID).Updates(Product{ProductName: product.ProductName, ProductQuantity: product.ProductQuantity, ProductPrice: product.ProductPrice, ProductImage: product.ProductImage, ProductUserID: product.ProductUserID}).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp := u.Message(true, "Update success")
	resp["product"] = product
	return resp
}