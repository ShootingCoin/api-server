package utils

import (
	"fmt"
)

// GetBucket converts price to bucket value
func GetBucket(price int64) (string, error) {
	if price < 0 {
		return "", fmt.Errorf("Invalid price: %d", price)
	}

	var bucket string
	switch {
	case price >= 100:
		bucket = "100"
	case price >= 50:
		bucket = "50"
	case price >= 10:
		bucket = "10"
	default:
		bucket = "1"
	}

	return bucket, nil
}
