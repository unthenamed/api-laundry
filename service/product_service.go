package service

import (
	"api-laundry/model"
	"errors"
)

func (s *laundryService) InsertProduct(mProduct model.Products) (model.Products, error) {
	return s.products.InsertProduct(mProduct)
}

func (s *laundryService) GetProductById(id int) (model.Products, error) {
	return s.products.GetProductById(id)
}

func (s *laundryService) GetAllProduct(productName string) ([]model.Products, error) {
	return s.products.GetAllProduct(productName)
}

func (s *laundryService) UpdateProductById(id int, mProduct model.Products) (model.Products, error) {
	oldProduct, err := s.products.GetProductById(id)
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

	return s.products.UpdateProductById(id, mProduct)
}

func (s *laundryService) DeleteProductById(id int) error {

	var err error
	_, err = s.products.GetProductById(id)
	if err != nil {
		return err
	}

	var trx []model.Response
	trx, err = s.transactions.GetAllTransaction(model.Transaction{})
	if err != nil {
		return err
	}

	for _, t := range trx {

		for _, b := range t.BillDetails {
			if b.Product.Id == id {
				return errors.New("constrain")
			}
		}
	}

	return s.products.DeleteProductById(id)
}
