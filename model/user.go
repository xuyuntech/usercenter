package model

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Gender     string `json:"gender"`
	Age        string `json:"age"`
	TotalSpend string `json:"totalSpend"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}
