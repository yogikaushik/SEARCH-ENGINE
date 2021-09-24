package controller

import (
	"net/http"
	"strings"

	"github.com/yogeshkaushik1904/search-engine/entity"
	"github.com/yogeshkaushik1904/search-engine/service"
)

var (
	svc service.OrderService
)

type controller struct{}

type OrderController interface {
	GetOrders(w http.ResponseWriter, r *http.Request)
	GetOrderById(w http.ResponseWriter, r *http.Request)
	SearchOrders(w http.ResponseWriter, r *http.Request)
	AddOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
}

func NewOrderController(service service.OrderService) OrderController {
	svc = service
	return &controller{}
}
func (*controller) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := svc.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = orders.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json "+err.Error(), http.StatusInternalServerError)
	}
}

//-------------------------------------------GetOrderById returns Order from ES for given Id-------------------------
func (*controller) GetOrderById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/orders/")
	orders, err := svc.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = orders.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

//----------------------SearchOrders returns all the Orders for given query and filters from ES-------------------------------------
func (*controller) SearchOrders(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	orders, err := svc.Search(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = orders.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json "+err.Error(), http.StatusInternalServerError)
	}
}

//-------------------------------------------AddOrder adds a new document in the ES----------------------------
func (*controller) AddOrder(w http.ResponseWriter, r *http.Request) {
	odr := entity.Order{}
	err := odr.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid order", http.StatusInternalServerError)
		return
	}
	data, err := svc.Add(&odr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//-------------------------------------DeleteOrder deletes a the document for given id in the ES----------------------------
func (*controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/orders/delete/")
	data, err := svc.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//-------------------------------------UpdateOrder updates a the document for given id in the ES----------------------------
func (*controller) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/orders/update/")
	odr := entity.Order{}
	err := odr.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid order", http.StatusInternalServerError)
		return
	}
	data, err := svc.Update(id, &odr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
