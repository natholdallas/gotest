package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Result struct {
	RFC3339     time.Time `json:"rfc3339"`
	RFC3339Nano time.Time `json:"rfc3339nano"`
	ISO8601     time.Time `json:"iso8601"`
	DateOnly    time.Time `json:"dateOnly"`
	Unix        time.Time `json:"unix"`
	UnixMs      time.Time `json:"unixMs"`
}

func main() {
	d, _ := os.ReadFile("./cmd/json-date/time.json")
	var result Result
	json.Unmarshal(d, &result)

	fmt.Printf("RFC3339: %v\n", result.RFC3339)
	fmt.Printf("RFC3339Nano: %v\n", result.RFC3339Nano)
	fmt.Printf("ISO8601: %v\n", result.ISO8601)
	fmt.Printf("DateOnly: %v\n", result.DateOnly)
	fmt.Printf("Unix: %v\n", result.Unix)
	fmt.Printf("UnixMs: %v\n", result.UnixMs)
}
