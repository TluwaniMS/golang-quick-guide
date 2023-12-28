package main
import ("fmt")

func main()  {
	number := 13
	
	switch number {
	case 1,3,5:
		fmt.Println("Grouped with the first three!")
	case 7,9,11:
		fmt.Println("Grouped with the middle three!")
	case 13,15,17:
		fmt.Println("Grouped with the last three")
	default:
		fmt.Println("Number provided is not available in any grouping!")
	}
}