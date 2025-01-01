package service

import (
	"api-laundry/model"
	"api-laundry/repo"
)

type TransactionService interface {
	InsertTransaction(model.Transaction) (model.Transaction, error)
	GetAllTransaction(model.Transaction) ([]model.Response, error)
	GetTransactionById(int) (model.Transaction, error)
}

type transactionService struct {
	repo repo.TransactionRepo
}

func (t *transactionService) InsertTransaction(mTransaction model.Transaction) (model.Transaction, error) {

	return t.repo.InsertTransaction(mTransaction)
}

func (t *transactionService) GetAllTransaction(mTransaction model.Transaction) ([]model.Response, error) {
	return t.repo.GetAllTransaction(mTransaction)
}

func (t *transactionService) GetTransactionById(id int) (model.Transaction, error) {
	return t.repo.GetTransactionById(id)
}

func ObjTransactionService(repo repo.TransactionRepo) TransactionService {
	return &transactionService{repo}
}
