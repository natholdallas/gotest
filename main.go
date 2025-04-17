package main

import t "github.com/natholdallas/gotest/pkg/tools"

type RateResponse struct {
	Result          string             `json:"result"`
	Documentation   string             `json:"documentation"`
	Terms           string             `json:"terms_of_use"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func main() {
	data := RateResponse{ConversionRates: make(map[string]float64)}

	t.PrintStruct(data)
}
