package main

import (
	"encoding/json"
	"fmt"
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
	d := []byte(`{
		"rfc3339": "2023-04-01T15:04:05Z",
		"rfc3339nano": "2023-04-01T15:04:05.999999999Z",
		"iso8601": "2023-04-01T15:04:05+08:00",
		"dateOnly": "2023-04-01",
		"unix": 1680354245,
		"unixMs": 1680354245123
	}`)

	var result Result

	json.Unmarshal(d, &result)
	fmt.Printf("RFC3339: %v\n", result.RFC3339)
	fmt.Printf("RFC3339Nano: %v\n", result.RFC3339Nano)
	fmt.Printf("ISO8601: %v\n", result.ISO8601)
	fmt.Printf("DateOnly: %v\n", result.DateOnly)
	fmt.Printf("Unix: %v\n", result.Unix)
	fmt.Printf("UnixMs: %v\n", result.UnixMs)
}
