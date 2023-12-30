package main
import ("fmt")

func main()  {
	fuelPrices := map[string]float32{"Unleaded":25.8,"Diesel":24.32,"LPG":58}

	fmt.Println("Map of fuel prices :",fuelPrices)
	fmt.Println("Cost of diesel :",fuelPrices["Diesel"])
}