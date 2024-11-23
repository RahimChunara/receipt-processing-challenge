package main

// receipt structure
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// receipt item
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// response after processing
type ProcessResponse struct {
	ID string `json:"id"`
}

// response on requesting receipt
type PointsResponse struct {
	Points int `json:"points"`
}
