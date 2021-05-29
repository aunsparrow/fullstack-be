package requestModels

type CreateProduct struct {
	ProductName   string  `json:"product_name"`
	ProductDetail string  `json:"product_detail"`
	ProductUnit   string  `json:"product_unit"`
	ProductPrice  float32 `json:"product_price"`
	CategoryId    string  `json:"category_id"`
	ShopId        string  `json:"shop_id"`
}

type UpdateProduct struct {
	ProductId     string  `json:"product_id"`
	ProductName   string  `json:"product_name"`
	ProductDetail string  `json:"product_detail"`
	ProductUnit   string  `json:"product_unit"`
	ProductPrice  float32 `json:"product_price"`
	CategoryId    string  `json:"category_id"`
}

type DeleteProduct struct {
	ProductId string `json:"product_id"`
}

type GetAllProduct struct {
	ProductId string `json:"product_id"`
}
