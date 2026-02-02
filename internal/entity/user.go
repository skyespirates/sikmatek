package entity

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	RoleId   int    `json:"role_id"`
}

type UserDetail struct {
	Id         int
	Email      string
	Password   string
	RoleId     int
	ConsumerId int
}

type RegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
