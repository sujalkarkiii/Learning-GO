package main

import (
	"cryptomaster/api"
	"fmt"
	"sync"

)

func main() {
currencies:=[]string{"BTC","ETH","BCH"}
var wg sync.WaitGroup
	for _,currency :=range currencies{
		wg.Add(1)
		go func (currencycode string)  {
			getCurrenctData(currencycode)
		wg.Done()
		}(currency)
	}
	wg.Wait()
}


func getCurrenctData(currency string)  {
	rate,err:=api.GetRate(currency)
		if err != nil {
		fmt.Println("BTC price:", rate.Price)

	}

}