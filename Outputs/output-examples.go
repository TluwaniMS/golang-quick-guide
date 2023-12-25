package main
import ("fmt")

func main()  {
	var name, age = "John", 18

	fmt.Println(name,age)

	fmt.Printf("name has a value: %v and type: %T\n", name, name)

	fmt.Printf("age has a value: %v and type: %T\n", age, age)
}