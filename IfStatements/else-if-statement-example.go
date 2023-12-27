package main
import ("fmt")

func main()  {
	number := 11
	
	if(number <= 5){
		fmt.Println("We're below 5!")
	}else if(number > 5 && number <= 10){
		fmt.Println("We're above 5!")
	}else{
		fmt.Println("We're way beyond 5!")
	}
}