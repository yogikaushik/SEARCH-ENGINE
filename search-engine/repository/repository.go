package repository

import (
	"net/url"

	"github.com/yogeshkaushik1904/search-engine/entity"
)

type OrderRepository interface {
	Add(order *entity.Order) ([]byte, error)
	Update(id string, order *entity.Order) ([]byte, error)
	Delete(id string) ([]byte, error)
	FindAll() (entity.Orders, error)
	FindById(Id string) (entity.Orders, error)
	Search(v url.Values) (entity.Orders, error)
}
