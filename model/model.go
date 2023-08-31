package model

// O serviço A seria um serviço de avaliação de produtos (review-service), uma avaliação é composta por:
// review_id, product_id, user_id, grade, created_at;

type Review struct {
	ReviewId  string `json:"review_id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Grade     int    `json:"grade"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
