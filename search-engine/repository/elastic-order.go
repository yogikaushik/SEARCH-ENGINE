package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/olivere/elastic"
	"github.com/yogeshkaushik1904/search-engine/entity"
)

type repo struct {
}

var (
	client *elastic.Client
)

func NewElasticOrderRepository() OrderRepository {
	c, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Printf("Error connecting with ES : %s\n", err)
		os.Exit(1)
	}
	client = c
	return &repo{}
}

func (r *repo) Add(o *entity.Order) ([]byte, error) {
	addResult, err := client.Index().
		Index("kibana_sample_data_ecommerce").
		BodyJson(&o).
		Do(context.Background())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errors.New("error adding Order")
	}
	data, err := json.Marshal(addResult)
	if err != nil {
		return nil, errors.New("unable to marshal result")
	}
	return data, nil
}

func (*repo) Update(id string, order *entity.Order) ([]byte, error) {
	updateResult, err := client.Update().
		Index("kibana_sample_data_ecommerce").
		Id(id).
		Doc(&order).
		Do(context.Background())

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errors.New("error occurred while updating order")
	}
	data, err := json.Marshal(updateResult)
	if err != nil {
		return nil, errors.New("unable to marshal update result")
	}
	return data, nil
}

func (*repo) Delete(id string) ([]byte, error) {
	deleteResult, err := client.Delete().
		Index("kibana_sample_data_ecommerce").
		Id(id).
		Do(context.Background())

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errors.New("error occurred while deleting order")
	}
	data, err := json.Marshal(deleteResult)
	if err != nil {
		return nil, errors.New("unable to marshal delete result")
	}
	return data, nil
}

func (r *repo) FindAll() (entity.Orders, error) {
	query := elastic.NewMatchAllQuery()
	getResult, err := client.Search().
		Index("kibana_sample_data_ecommerce").
		Query(query).
		From(0).Size(100).
		Do(context.Background())

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errors.New("error occurred while fetching orders")
	}
	orders, err := entity.GetOrdersList(getResult)
	if err != nil {
		return nil, errors.New("unable to get orders list")
	}
	return orders, nil
}

func (*repo) FindById(id string) (entity.Orders, error) {

	getResult, err := client.Get().
		Index("kibana_sample_data_ecommerce").
		Id(id).
		Do(context.Background())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errors.New("error occurred while fetching orders")
	}
	orders := entity.Orders{}
	odr := entity.Order{}
	err = json.Unmarshal(getResult.Source, &odr)
	if err != nil {
		return nil, errors.New("unable to parse order")
	}
	orders = append(orders, &odr)
	return orders, nil
}

func (*repo) Search(v url.Values) (entity.Orders, error) {

	filters := entity.GetFilters(v)
	query := elastic.NewMultiMatchQuery(strings.ToLower(v["query"][0]), "products.product_name", "category")
	bf := elastic.NewBoolQuery().Must(query).Filter(filters...)

	searchResult, err := client.Search().
		Index("kibana_sample_data_ecommerce").
		Query(bf).
		From(0).Size(100).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errors.New("error occurred while fetching orders")
	}
	orders, err := entity.GetOrdersList(searchResult)
	if err != nil {
		return nil, errors.New("unable to get orders list")
	}
	return orders, nil
}
