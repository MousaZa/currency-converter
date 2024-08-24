package main

import (
	"fmt"
	"github.com/MousaZa/currency-converter/api"
	"github.com/charmbracelet/huh"
	"log"
	"strconv"
)

func main() {
	var (
		base   string
		dest   string
		amount string
	)
	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[string]().
				Title("Choose the base currency").
				Options(
					huh.NewOption("Turkish Lira", "try"),
					huh.NewOption("American Dollar", "usd"),
					huh.NewOption("Euro", "eur"),
				).
				Value(&base), // store the chosen option in the "burger" variable

			// Let the user select multiple toppings.
			huh.NewSelect[string]().
				Title("Choose the destination currency").
				Options(
					huh.NewOption("Turkish Lira", "try"),
					huh.NewOption("American Dollar", "usd"),
					huh.NewOption("Euro", "eur"),
				).
				Value(&dest),
			huh.NewInput().
				Title("The amount").
				Value(&amount).
				Validate(func(str string) error {
					_, err := strconv.ParseFloat(amount, 64)
					return err
				}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	rate, err := api.GetPrices(base, dest)
	amountf, err := strconv.ParseFloat(amount, 64)

	result := amountf * rate

	fmt.Printf("%v %v is equal %.2f %v\n", amountf, base, result, dest)
}
