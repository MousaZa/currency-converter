package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetPrices(base, dest string) (float64, error) {
	url := "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.json"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err", err)
		return -1, err
	}
	if res.StatusCode == http.StatusOK {
		result := Data{}
		body, _ := io.ReadAll(res.Body)
		err := json.Unmarshal(
			body,
			&result,
		)
		if err != nil {
			fmt.Println("err", err)
			return -1, err
		}
		rate := result.Rate[dest] / result.Rate[base]
		return rate, nil
	}
	return -1, errors.New("error")
}

type Data struct {
	Rate map[string]float64 `json:"eur"`
	Date string             `json:"date"`
}
