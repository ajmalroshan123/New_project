package db

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Credentials
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Favourite struct {
	UserID       string `json:"user_id"`
	RestaurantID string `json:"restaurant_id"`
}
