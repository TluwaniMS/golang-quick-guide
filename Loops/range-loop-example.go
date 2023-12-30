package main
import ("fmt")

func main()  {
	names := [5]string{"Thabo","Leago","Phuluphelo","Leeba","Morokolo"}

	for index,value := range names{
		fmt.Printf("%v\t%v\n", index, value)
	}
}