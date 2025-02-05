package service

import (
	"api-laundry/model"
	"api-laundry/repo"
)

type ProductService interface {
	InsertProduct(model.Products) (model.Products, error)
	GetProductById(int) (model.Products, error)
	GetAllProduct(string) ([]model.Products, error)
	UpdateProductById(int, model.Products) (model.Products, error)
	DeleteProductById(int) error
}

type productService struct {
	Repo repo.ProductRepo
}

func (p *productService) InsertProduct(mProduct model.Products) (model.Products, error) {
	return p.Repo.InsertProduct(mProduct)
}

func (p *productService) GetProductById(id int) (model.Products, error) {
	return p.Repo.GetProductById(id)
}

func (p *productService) GetAllProduct(productName string) ([]model.Products, error) {
	return p.Repo.GetAllProduct(productName)
}

func (p *productService) UpdateProductById(id int, mProduct model.Products) (model.Products, error) {
	oldProduct, err := p.Repo.GetProductById(id)
	if err != nil {
		return model.Products{}, err
	}

	if mProduct.Name == "" {
		mProduct.Name = oldProduct.Name
	}
	if mProduct.Price == 0 {
		mProduct.Price = oldProduct.Price
	}
	if mProduct.Unit == "" {
		mProduct.Unit = oldProduct.Unit
	}

	return p.Repo.UpdateProductById(id, mProduct)
}

func (p *productService) DeleteProductById(id int) error {

	return p.Repo.DeleteProductById(id)
}

// ObjProductService is a function to create a new object of ProductService
func ObjProductService(repo repo.ProductRepo) ProductService {
	return &productService{Repo: repo}
}
