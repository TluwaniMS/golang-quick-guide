package main
import ("fmt")

func singleValueReturn(x int,y int)int{
	return x + y
}

func main()  {
	fmt.Println(singleValueReturn(2,5))
}