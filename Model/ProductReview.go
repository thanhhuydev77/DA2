package Model

type ProductReview struct {
	Id          	string `json:"id"`
	Name        	string	`json:"name"`
	ReviewUsername  string `json: "reviews.username"`
	Rating      	int	`json:"	reviews.rating"`
}

type ProductReviewUserBased struct {
	
}