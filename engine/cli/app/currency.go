package app

import (
	"engine/lib/structs"
	"fmt"
)

func (app *App) AppendContract(name, symbol, contractAddress string) {
	//structs.Currency{}

}

func (app *App) ActivateCurrency(symbol string) error {
	currency, err := app.srvc_c.Activate(app.ctx, &structs.Query_Currency{Symbol: symbol})
	if err != nil {
		return fmt.Errorf("ActivateCurrency: %s", err)
	}

	fmt.Println("Currency activated")
	fmt.Println(currency.Name)
	return nil
}

func (app *App) DeactivateCurrency(symbol string) error {
	currency, err := app.srvc_c.Deactivate(app.ctx, &structs.Query_Currency{Symbol: symbol})
	if err != nil {
		return fmt.Errorf("DeactivateCurrency: %s", err)
	}

	fmt.Println("Currency deactivated")
	fmt.Println(currency.Name)
	return nil
}
