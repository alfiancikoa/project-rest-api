package databases

import (
	"project1/config"
	"project1/models"
)

func Insert(product models.Product) (models.Product, error) {
	if err := config.DB.Save(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetAll() ([]models.ListProductRespon, error) {
	products := []models.ListProductRespon{}
	query := config.DB.Table("products").Select(
		"products.id, products.name, products.price, products.description, products.qty, products.created_at").
		Where("products.deleted_at is NULL").Find(&products)
	if query.Error != nil {
		return nil, query.Error
	}
	return products, nil
}

func GetFilterByName(request string) ([]models.ListProductRespon, error) {
	products := []models.ListProductRespon{}
	orderBy := "products.name"
	if request != "ascending" {
		orderBy = "products.name DESC"
	}
	query := config.DB.Table("products").Select(
		"products.id, products.name, products.price, products.description, products.qty, products.created_at").
		Where("products.deleted_at is NULL").Order(orderBy).Find(&products)
	if query.Error != nil {
		return nil, query.Error
	}
	return products, nil
}

func GetFilterPrice(request string) ([]models.ListProductRespon, error) {
	products := []models.ListProductRespon{}
	orderBy := "products.price"
	if request != "lower" {
		orderBy = "products.price DESC"
	}
	query := config.DB.Table("products").Select(
		"products.id, products.name, products.price, products.description, products.qty, products.created_at").
		Where("products.deleted_at is NULL").Order(orderBy).Find(&products)
	if query.Error != nil {
		return nil, query.Error
	}
	return products, nil
}

func GetFilterNewProduct() ([]models.ListProductRespon, error) {
	products := []models.ListProductRespon{}
	query := config.DB.Table("products").Select(
		"products.id, products.name, products.price, products.description, products.qty, products.created_at").
		Where("products.deleted_at is NULL").Order("products.id DESC").Find(&products)
	if query.Error != nil {
		return nil, query.Error
	}
	return products, nil
}

// Fungsi untuk mengecek ketersediaan data di database
func CheckDatabase(coloumn string, data string) (int64, error) {
	product := models.Product{}
	tx := config.DB.Where(coloumn+"=?", data).Find(&product)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return tx.RowsAffected, nil
}
