package service

import (
	"errors"
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"
)

type SaleService interface {
	CreateSale(sale *model.Sale, items []model.SaleItem) error
	GetAll() ([]model.Sale, error)
	GetByID(id int64) (*model.Sale, error)
	GetSalesReport() (*model.SalesReport, error)
}

type saleService struct {
	repo repository.SaleRepository
}

func NewSaleServer(r repository.SaleRepository) SaleService {
	return &saleService{repo: r}
}

func (s *saleService) CreateSale(sale *model.Sale, items []model.SaleItem) error {
	if len(items) == 0 {
		return errors.New("data item tidak boleh kosong")
	}

	var total int64 = 0
	for _, item := range items {
		if item.Quantity <= 0 || item.Price <= 0 {
			return errors.New("quantity dan price lebih dari 0")
		}
		total += item.Quantity * item.Price
	}
	sale.Total = total

	return  s.repo.CreateSale(sale, items)
}

func (s *saleService) GetAll() ([]model.Sale, error) {
	return s.repo.GetAll()
}

func (s *saleService) GetByID(id int64) (*model.Sale, error) {
	return s.repo.GetByID(id)
}

func (s *saleService) GetSalesReport() (*model.SalesReport, error) {
	return s.repo.GetSalesReport()
}