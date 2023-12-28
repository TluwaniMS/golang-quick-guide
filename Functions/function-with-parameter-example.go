package main
import ("fmt")

func greetTheWorldNamely(name string,surname string){
	fmt.Println("Hello",name,surname,"from the world of Go!!!")
}

func main()  {
	greetTheWorldNamely("John","Doe")
}