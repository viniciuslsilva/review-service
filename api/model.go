package api

type ReviewPayload struct {
	ReviewId  string `json:"review_id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Grade     int    `validate:"min=0,max=5" json:"grade"`
}
