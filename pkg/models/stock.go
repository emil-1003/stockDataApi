package models

type StockData struct {
	Symbol 	string 	`json:"symbol"`
	Price 	string 	`json:"price"`
	Change 	string 	`json:"change"`
}