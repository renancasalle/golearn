package service 

import (
	"golearn/model"
)

type ProductService struct {

}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (ps *ProductService) GetProducts() ([]model.Product, error){
	return []model.Product{}, nil
}