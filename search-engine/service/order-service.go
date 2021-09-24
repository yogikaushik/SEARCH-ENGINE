package service

import (
	"errors"
	"io"
	"net/url"

	"github.com/yogeshkaushik1904/search-engine/entity"
	"github.com/yogeshkaushik1904/search-engine/repository"
)

type OrderService interface {
	MiddlewareValidateOrder(io.Reader) error
	Add(order *entity.Order) ([]byte, error)
	Update(Id string, order *entity.Order) ([]byte, error)
	Delete(Id string) ([]byte, error)
	FindAll() (entity.Orders, error)
	FindById(Id string) (entity.Orders, error)
	Search(v url.Values) (entity.Orders, error)
}

type service struct{}

var (
	repo repository.OrderRepository
)

func NewOrderService(repository repository.OrderRepository) OrderService {
	repo = repository
	return &service{}
}

func (*service) MiddlewareValidateOrder(r io.Reader) error {
	odr := entity.Order{}
	err := odr.FromJSON(r)
	if err != nil {
		return errors.New("error reading order")
	}
	return nil
}

func (*service) Add(order *entity.Order) ([]byte, error) {
	return repo.Add(order)
}
func (*service) Update(Id string, order *entity.Order) ([]byte, error) {
	return repo.Update(Id, order)
}
func (*service) Delete(Id string) ([]byte, error) {
	return repo.Delete(Id)
}
func (*service) FindAll() (entity.Orders, error) {
	return repo.FindAll()
}
func (*service) FindById(Id string) (entity.Orders, error) {
	return repo.FindById(Id)
}

func (*service) Search(v url.Values) (entity.Orders, error) {
	return repo.Search(v)
}
