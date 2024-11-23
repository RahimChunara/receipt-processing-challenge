package main

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func calculatePoints(receipt Receipt) int {
	points := 0

	// one point for every alphanumeric character in the retailer name.
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			points++
		}
	}

	// 50 points if the total is a round dollar amount with no cents.
	if isRoundDollarAmount(receipt.Total) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// multiply price by 0.2 and round up if description length is multiple of 3.
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		descLen := len(desc)
		if descLen%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil || price < 0 {
				continue 
			}
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
		}
	}

	// 6 points if the day in the purchase date is odd.
	if isDayOdd(receipt.PurchaseDate) {
		points += 6
	}

	// 10 points if the time is between 2:00pm and 4:00pm.
	if isBetweenTwoAndFourPM(receipt.PurchaseTime) {
		points += 10
	}

	return points
}

func isRoundDollarAmount(total string) bool {
	amount, err := strconv.ParseFloat(total, 64)
	if err != nil || amount < 0 {
		return false
	}
	cents := int(math.Round(amount * 100))
	return cents%100 == 0
}

func isMultipleOfQuarter(total string) bool {
	amount, err := strconv.ParseFloat(total, 64)
	if err != nil || amount < 0 {
		return false
	}
	cents := int(math.Round(amount * 100))
	return cents%25 == 0
}

func isDayOdd(purchaseDate string) bool {
	date, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return false
	}
	return date.Day()%2 == 1
}

func isBetweenTwoAndFourPM(purchaseTime string) bool {
	t, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return false
	}
	start, _ := time.Parse("15:04", "14:00")
	end, _ := time.Parse("15:04", "16:00")
	return !t.Before(start) && t.Before(end)
}

func isValidReceipt(receipt Receipt) bool {
	// checks for missing or empty required fields
	if strings.TrimSpace(receipt.Retailer) == "" ||
		strings.TrimSpace(receipt.PurchaseDate) == "" ||
		strings.TrimSpace(receipt.PurchaseTime) == "" ||
		strings.TrimSpace(receipt.Total) == "" ||
		len(receipt.Items) == 0 {
		return false
	}

	// validate total is a valid, non-negative number
	totalAmount, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil || totalAmount < 0 {
		return false
	}

	// validate each item
	for _, item := range receipt.Items {
		if strings.TrimSpace(item.ShortDescription) == "" {
			return false
		}
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil || price < 0 {
			return false
		}
	}

	// validates date and time formats
	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return false
	}
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return false
	}

	return true
}
