package repository

//go:generate mockgen -destination=../../mocks/company_repo.go -package=mocks github.com/wgarunap/xm-rest-api/domain/repository Company

import "github.com/wgarunap/xm-rest-api/domain"

type Company interface {
	Create(company domain.Company) error
	Update(company domain.Company) error
	Get(filters ...Filter) ([]domain.Company, error)
	Delete(filters ...Filter) error
}


type Filter struct {
	FieldName string
	Value     interface{}
}
