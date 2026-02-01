package entity

type User struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	RoleId     int    `json:"role_id"`
	ConsumerId int    `json:"consumer_id"`
}

type RegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
