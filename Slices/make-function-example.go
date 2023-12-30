package main
import ("fmt")

func main()  {
	slice := make([]int,5,10)

	fmt.Println("The slice length :",len(slice))
	fmt.Println("The slice capacity :",cap(slice))
	fmt.Println("The slice itself :",slice)
}