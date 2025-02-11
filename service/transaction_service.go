package service

import (
	"api-laundry/model"
)

func (s *laundryService) InsertTransaction(mTransaction model.Transaction) (model.Response, error) {

	return s.transactions.InsertTransaction(mTransaction)
}

func (s *laundryService) GetAllTransaction(mTransaction model.Transaction) ([]model.Response, error) {
	mTransaction.QueryDate()
	return s.transactions.GetAllTransaction(mTransaction)
}

func (s *laundryService) GetTransactionById(id int) (model.Transaction, error) {
	return s.transactions.GetTransactionById(id)
}
