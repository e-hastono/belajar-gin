package main

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	database.StartDB()
	db = database.GetDB()
	// createUser("johndoe@gmail.com")
	// getUserById(1)
	// updateUserById(1, "johndoe@outlook.com")
	// createProduct(1, "YLO", "YYYY")
	// getUsersWithProducts()
	deleteProductById(1)
}

func createUser(email string) {
	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}

func getUserById(id uint) {
	user := models.User{}

	err := db.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v \n", user)
}

func updateUserById(id uint, email string) {
	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func createProduct(uid uint, brand string, name string) {
	product := models.Product{
		UserID: uid,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("New Product Data:", product)
}

func getUsersWithProducts() {
	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error geting user data with products: ", err.Error())
		return
	}

	fmt.Println("User Data with Products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted ", id)
}
