package request

type ProductRequest struct {
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}
