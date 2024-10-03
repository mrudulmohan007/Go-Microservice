package models

type Product struct {
	ProductID    string `json:"productid"`
	ProductName  string `json:"productname"`
	ProductPrice int    `json:"productprice"`
	UserID       string `json:"userId"`
}

// type Product struct {
// 	ProductID    string `json:"productid"`
// 	ProductName  string `json:"productname"`
// 	ProductPrice int    `json:"productprice"`
// 	UserID       string `json:"userId"`
// }
