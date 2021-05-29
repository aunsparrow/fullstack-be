package requestModels

type CreateShop struct {
	ShopName string `json:"shop_name"`
	Tel      string `json:"tel"`
	Address  string `json:"address"`
}

type UpdateShop struct {
	ShopId   string `json:"shop_id"`
	ShopName string `json:"shop_name"`
	Tel      string `json:"tel"`
	Address  string `json:"address"`
}

type DeleteShop struct {
	ShopId string `json:"shop_id"`
}

type GetAllShop struct {
	ShopId string `json:"shop_id"`
}

type Pagination struct {
	FilterText string `json:"filter_text"`
	PageNumber int    `json:"page_number"`
	PageSize   int    `json:"page_size"`
}
