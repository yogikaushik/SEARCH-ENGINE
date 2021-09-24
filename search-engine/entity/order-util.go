package entity

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/olivere/elastic"
)

func (o *Order) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(o)
}

func (o *Orders) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}
func GetOrdersList(esr *elastic.SearchResult) (odrs Orders, err error) {
	for _, hit := range esr.Hits.Hits {
		o := Order{}
		err = json.Unmarshal(hit.Source, &o)
		if err != nil {
			return nil, err
		}
		odrs = append(odrs, &o)
	}
	return odrs, err
}

func GetFilters(v map[string][]string) (filters []elastic.Query) {
	for k, v := range v {
		switch k {
		case "category":
			cf := elastic.NewTermQuery("category", strings.ToLower(v[0]))
			filters = append(filters, cf)
		case "country":
			ctf := elastic.NewTermQuery("geoip.country_iso_code", v[0])
			filters = append(filters, ctf)
		case "manufacturer":
			mf := elastic.NewTermQuery("manufacturer", strings.ToLower(v[0]))
			filters = append(filters, mf)
		case "minprice":
			mnpf := elastic.NewRangeQuery("taxful_total_price").From(v[0])
			filters = append(filters, mnpf)
		case "maxprice":
			mxpf := elastic.NewRangeQuery("taxful_total_price").To(v[0])
			filters = append(filters, mxpf)
		default: //do nothing
		}
	}

	return filters
}
