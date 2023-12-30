package main
import ("fmt")

func main()  {
	slice1 := []int{1,2,3,4,5}

	combinedSlice := append(slice1,6,7)

	fmt.Println("The combined slice :",combinedSlice)
}