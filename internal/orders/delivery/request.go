package delivery

type RequestCheckout struct {
	UserValue string `json:"user_value"`
	ProductId uint   `json:"product_id"`
}
