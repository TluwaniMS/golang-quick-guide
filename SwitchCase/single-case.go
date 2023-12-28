package main
import ("fmt")

func main()  {
	number := 11
	
	switch number {
	case 1:
		fmt.Println("It's the first number!")
	case 9:
		fmt.Println("It's the second number!")
	case 11:
		fmt.Println("It's the third number!")
	default:
		fmt.Println("The number is not provided for!")
	}
}