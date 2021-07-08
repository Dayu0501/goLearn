package services

import "context"

type ProdService struct {
	prodList []string
}

func (p *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {

	return &ProdResponse{ProdStock: 20}, nil
}

func (p *ProdService) GetProdName(context.Context, *ProdRequestName) (*ProdResponseName, error) {

	return &ProdResponseName{Name: "apple"}, nil
}
