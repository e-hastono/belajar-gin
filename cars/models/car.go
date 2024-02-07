package models

// Car represents model of a car
type Car struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Brand string `gorm:"size:255;not null;index" json:"brand"`
	Model string `gorm:"size:255;not null;index" json:"model"`
	Price int    `gorm:"type:int" json:"price"`
}

func (newCar *Car) CreateCar() (*Car, error) {
	db := GetDB()

	if err := db.Create(&newCar).Error; err != nil {
		return &Car{}, err
	}

	return newCar, nil
}

func (car *Car) UpdateCar() (*Car, error) {
	db := GetDB()
	carID := car.ID
	var updatedCar Car

	err := db.Model(&updatedCar).Where("id = ?", carID).Updates(Car{Brand: car.Brand, Model: car.Model, Price: car.Price}).Error

	if err != nil {
		return &Car{}, err
	}

	return &updatedCar, nil
}

func GetAllCars() ([]Car, error) {
	db := GetDB()
	var cars []Car
	err := db.Find(&cars).Error

	if err != nil {
		return []Car{}, err
	}

	return cars, err
}

func GetCar(carID int) (Car, error) {
	db := GetDB()
	var car Car

	if err := db.First(&car, carID).Error; err != nil {
		return car, err
	}

	return car, nil
}

func DeleteCar(carID int) (int, error) {
	db := GetDB()

	if err := db.Where("id = ?", carID).Delete(&Car{}, carID).Error; err != nil {
		return carID, err
	}

	return carID, nil
}
