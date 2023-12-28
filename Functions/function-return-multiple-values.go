package main
import ("fmt")

func multipleValueReturn(name string,age string)(nameGreeting string,ageShout string){
	nameGreeting = "Hello, hello " + name + "..."
	ageShout = "You're looking great, are you really " + age + " ?"
	return
}

func main()  {
	nameGreeting,ageShout:= multipleValueReturn("John","45")

	fmt.Println(nameGreeting)
	fmt.Println(ageShout)
}