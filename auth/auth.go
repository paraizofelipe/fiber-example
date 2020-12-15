package auth

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type UserData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
