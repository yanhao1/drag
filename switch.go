package main

import (
	"fmt"
	"time"
)

func main()  {
	i :=2
	fmt.Print("write",i,"as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("there")
	}
	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("it is weeked!")
	default:
		fmt.Println("It's a weeked!")
	}

	t :=time.Now()
	switch {
	case t.Hour()<12:
		fmt.Println("It is befor noon")
	default:
		fmt.Println("it is after noon")
	}

	whatAmI :=func(i interface{}){
		switch t:=i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Don't know type %T\n",t)

		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}