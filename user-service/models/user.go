package models

type User struct {
	UserID    string `json:"userid"`
	UserName  string `json:"username"`
	UserEmail string `json:"useremail"`
}
