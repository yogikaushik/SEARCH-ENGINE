package entity

type Order struct {
	ID            int       `json:"order_id"`
	OrderDate     string    `json:"order_date"`
	TotalQuantity int       `json:"total_quantity"`
	Price         float64   `json:"taxful_total_price"`
	Category      []string  `json:"category"`
	Customer      string    `json:"customer_full_name"`
	CustomerEmail string    `json:"email"`
	CustomerPhone string    `json:"customer_phone"`
	Manufacturer  []string  `json:"manufacturer"`
	Products      []Product `json:"products"`
	GeoIP         GeoIP     `json:"geoip"`
	SKU           []string  `json:"sku"`
}

type Product struct {
	ID          int     `json:"product_id"`
	Name        string  `json:"product_name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Discount    float32 `json:"discount_percentage"`
	Tax         float32 `json:"tax_amount"`
	UnitPrice   float32 `json:"base_unit_price"`
	TotalPrice  float32 `json:"taxful_price"`
}

type GeoIP struct {
	City      string `json:"city_name"`
	Country   string `json:"country_iso_code"`
	Region    string `json:"region_name"`
	Continent string `json:"continent_name"`
}

type Orders []*Order
