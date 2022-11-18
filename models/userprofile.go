package models

type UserProfile struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Country  string `json:"country"`
}